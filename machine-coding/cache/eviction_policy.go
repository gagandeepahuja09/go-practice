package cache

import (
	"container/list"
	"sync"
)

type EvictionPolicy interface {
	evict(cm *CacheManager)
	update(key string)
}

type LruEvictionPolicy struct {
	maxSize      int
	queue        *list.List
	mutex        sync.Mutex
	elementIndex map[string]*list.Element
}

func (lru *LruEvictionPolicy) evict(cm *CacheManager) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if lru.queue.Len() >= lru.maxSize {
		oldest := lru.queue.Front()
		if oldest != nil {
			key := oldest.Value.(string)
			delete(cm.cache, key)
			delete(lru.elementIndex, key)
			lru.queue.Remove(oldest)
		}
	}
}

// update moves the accessed key to the end of the LRU queue if existing key.
// If it is a new key, it inserts are the end of the queue
// Optimized to O(1) by keeping pointer references of each key.
func (lru *LruEvictionPolicy) update(key string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if element, exists := lru.elementIndex[key]; exists {
		lru.queue.MoveToBack(element)
	} else {
		element := lru.queue.PushBack(key)
		lru.elementIndex[key] = element
	}
}

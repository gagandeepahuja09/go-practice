package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// note: while we could have used map[string]interface as well as value in the CacheManager
// struct instead of having values as CacheEntry. For production application, it would make sense
// to also store some metadata values as well.
// These can be useful for implementing cache eviction policies, supporting TTLs, cache stats
// and cache monitoring
type CacheEntry struct {
	value          interface{}
	creationTime   time.Time
	expirationTime time.Time
	accessCount    int
}

type CacheManager struct {
	cache          map[string]*CacheEntry
	evictionPolicy EvictionPolicy
	mutex          sync.Mutex
}

var instance *CacheManager
var once sync.Once

func GetCacheManagerInstance() *CacheManager {
	once.Do(func() {
		instance = &CacheManager{
			cache:          make(map[string]*CacheEntry),
			evictionPolicy: &LruEvictionPolicy{maxSize: 3, queue: list.New()},
		}
	})
	return instance
}

func (cm *CacheManager) Put(key string, value interface{}, ttlInSeconds int) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if val, ok := cm.cache[key]; ok {
		val.accessCount++
		cm.cache[key] = val
	} else {
		cm.cache[key] = &CacheEntry{
			value:          value,
			creationTime:   time.Now(),
			expirationTime: time.Now().Add(time.Duration(ttlInSeconds) * time.Second),
			accessCount:    0,
		}
	}

	cm.evictionPolicy.update(key)
	cm.evictionPolicy.evict(cm)
}

func (cm *CacheManager) Get(key string) (interface{}, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if cacheEntry, ok := cm.cache[key]; ok {
		if cacheEntry.expirationTime.Before(time.Now()) {
			cm.evictionPolicy.evict(cm)
		}
		return cacheEntry.value, nil
	}
	return nil, fmt.Errorf("key not found in cache")
}

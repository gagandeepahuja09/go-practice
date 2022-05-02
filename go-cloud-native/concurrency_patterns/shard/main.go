package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type Shard struct {
	sync.RWMutex
	m map[string]interface{}
}

type ShardedMap []*Shard

func NewShardedMap(nshards int) ShardedMap {
	shards := make([]*Shard, nshards)

	for i := 0; i < nshards; i++ {
		shard := make(map[string]interface{})
		shards[i] = &Shard{m: shard}
	}

	return shards
}

func (m ShardedMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[17]) // Pick an arbitrary byte as the hash
	// Since we are using a byte-sized value as the hash value, it can only handle upto
	// 255 shards. We can sprinkle some binary arithmetic, if we need more than that
	// eg. hash := int(sum[13]) << 8 | int(sum[17])
	return hash % len(m)
}

func (m ShardedMap) getShard(key string) *Shard {
	index := m.getShardIndex(key)
	return m[index]
}

func (m ShardedMap) Get(key string) interface{} {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	return shard.m[key]
}

func (m ShardedMap) Set(key string, val interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = val
}

func (m ShardedMap) Keys() []string {
	var wg sync.WaitGroup
	wg.Add(len(m))
	mutex := sync.Mutex{}
	var keys []string
	// concurrently getting all the keys
	for _, shard := range m {
		go func(s *Shard) {
			s.RLock()

			for val := range s.m {
				mutex.Lock()
				keys = append(keys, val)
				mutex.Unlock()
			}

			s.RUnlock()
			wg.Done()
		}(shard)
	}

	wg.Wait()
	return keys
}

func main() {
	shardedMap := NewShardedMap(7)
	shardedMap.Set("alpha", 1)
	shardedMap.Set("beta", 2)
	shardedMap.Set("gamma", 3)

	fmt.Println(shardedMap.Get("alpha"))
	fmt.Println(shardedMap.Get("beta"))
	fmt.Println(shardedMap.Get("gamma"))

	for _, k := range shardedMap.Keys() {
		fmt.Println(k)
	}
}

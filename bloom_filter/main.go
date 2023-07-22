package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

var mHasher hash.Hash32

func initMurmurHash() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func murmurHash(key string) int {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32()
	mHasher.Reset()
	return int(result)
}

type BloomFilter struct {
	filter []bool
	size   int
}

func NewBloomFilter(size int) *BloomFilter {
	initMurmurHash()
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func (b *BloomFilter) Add(key string) {
	idx := murmurHash(key) % b.size
	b.filter[idx] = true
}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurHash(key) % b.size
	return b.filter[idx]
}

func main() {
	bloom := NewBloomFilter(10)
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		bloom.Add(key)
		fmt.Printf("exists key %s %v\n", key, bloom.Exists(key))
	}

	fmt.Printf("exists key %s %v\n", "d", bloom.Exists("d"))
}

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
	filter []int8
	size   int
}

func NewBloomFilter(size int) *BloomFilter {
	initMurmurHash()
	return &BloomFilter{
		filter: make([]int8, size),
		size:   size,
	}
}

func (b *BloomFilter) Add(key string) {
	idx := murmurHash(key) % b.size
	arrIdx := idx / 8
	bitIdx := idx % 8
	b.filter[arrIdx] = b.filter[arrIdx] | (1 << bitIdx)
}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurHash(key) % b.size
	arrIdx := idx / 8
	bitIdx := idx % 8
	return b.filter[arrIdx]&(1<<bitIdx) > 0
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

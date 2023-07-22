package main

import (
	"fmt"
	"hash"
	"math/rand"

	"github.com/spaolacci/murmur3"
)

func murmurHash(mHash hash.Hash32, key string) int {
	mHash.Write([]byte(key))
	result := mHash.Sum32()
	mHash.Reset()
	return int(result)
}

type BloomFilter struct {
	filter  []int8
	size    int
	hashFns []hash.Hash32
}

func NewBloomFilter(size, numHashFns int) *BloomFilter {
	return &BloomFilter{
		filter:  make([]int8, size),
		size:    size,
		hashFns: initMurmurHash(numHashFns),
	}
}

func initMurmurHash(numHashFns int) []hash.Hash32 {
	var mHashers []hash.Hash32
	for i := 0; i < numHashFns; i++ {
		mHashers = append(mHashers, murmur3.New32WithSeed(rand.Uint32()))
	}
	return mHashers
}

func (b *BloomFilter) Add(key string) {
	for _, mHash := range b.hashFns {
		idx := murmurHash(mHash, key) % b.size
		arrIdx := idx / 8
		bitIdx := idx % 8
		b.filter[arrIdx] = b.filter[arrIdx] | (1 << bitIdx)
	}
}

func (b *BloomFilter) Exists(key string) bool {
	for _, mHash := range b.hashFns {
		idx := murmurHash(mHash, key) % b.size
		arrIdx := idx / 8
		bitIdx := idx % 8
		if b.filter[arrIdx]&(1<<bitIdx) == 0 {
			return false
		}
	}
	return true
}

func main() {
	bloom := NewBloomFilter(10, 1)
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		bloom.Add(key)
		fmt.Printf("exists key %s %v\n", key, bloom.Exists(key))
	}

	fmt.Printf("exists key %s %v\n", "d", bloom.Exists("d"))
}

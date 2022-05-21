package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

var cache *lru.Cache

func init() {
	cache, _ = lru.NewWithEvict(2, func(key, value interface{}) {
		fmt.Printf("Evicted: key=%v value=%v\n", key, value)
	})
}

func main() {
	cache.Add(1, "a")
	cache.Add(2, "b")

	fmt.Println(cache.Get(1)) // 1 now becomes MRU.

	cache.Add(3, "c")

	fmt.Println(cache.Get(2))
}

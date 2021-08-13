package main

import "fmt"

type lru struct {
}

func (_ *lru) evict(_ *cache) {
	fmt.Println("Evicted using lru")
}

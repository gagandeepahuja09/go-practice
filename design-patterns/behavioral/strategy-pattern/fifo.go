package main

import "fmt"

type fifo struct {
}

func (_ *fifo) evict(_ *cache) {
	fmt.Println("Evicted using fifo")
}

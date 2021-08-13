package main

import "fmt"

type lfu struct {
}

func (_ *lfu) evict(_ *cache) {
	fmt.Println("Evicted using lfu")
}

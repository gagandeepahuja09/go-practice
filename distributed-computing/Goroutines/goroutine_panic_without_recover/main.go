package main

import (
	"fmt"
	"sync"
)

func simpleFunc(index int, wg *sync.WaitGroup) {
	fmt.Println("Attempting x / (x - 10) where x = ", index, "answer is ",
		index/(index-10))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(40)
	for i := 0; i < 40; i += 1 {
		go simpleFunc(i, &wg)
	}
	wg.Wait()
}

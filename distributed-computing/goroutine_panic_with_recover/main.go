package main

import (
	"fmt"
	"sync"
)

func simpleFunc(index int, wg *sync.WaitGroup) {
	// Executing a call to recover inside a deferred function stops the
	// panicking sequence by restoring the normal execution and retrieves the error
	// value passed to the calling function.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from ", r)
		}
	}()
	// we have changed the order of wg.Done as it should be called even if the following
	// line fails.
	defer wg.Done()
	fmt.Println("Attempting x / (x - 10) where x = ", index, "answer is ",
		index/(index-10))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(40)
	for i := 0; i < 40; i += 1 {
		go simpleFunc(i, &wg)
	}
	wg.Wait()
}

package main

import (
	"flag"
	"fmt"
	"sync"
)

// waiting for our goroutines to finish

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		// increase a counter in waitGroup variable
		waitGroup.Add(1)
		go func(x int) {
			// when the goroutine finishes its job, call waitGroup.Done()
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// one of the values in state1 holds the counter. which increases and decreases
	// if the count of add and done don't match, then we'll get an error.
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting....")
}

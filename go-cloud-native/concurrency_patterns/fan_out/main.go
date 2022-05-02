package main

import (
	"fmt"
	"sync"
)

func Split(source <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)
	for i := 0; i < n; i++ {
		ch := make(chan int)
		dests = append(dests, ch)

		go func() {
			defer close(ch)

			// Each channel gets a separate goroutine that competes for reads
			for val := range source {
				ch <- val
			}
		}()
	}
	return dests
}

func main() {
	source := make(chan int)
	dests := Split(source, 5)

	// Send the numbers 1...10 to source and close it when we are done.
	go func() {
		for i := 1; i <= 10; i++ {
			source <- i
		}

		close(source)
	}()

	// Use waitgroup to wait until all output channels close.
	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, ch := range dests {
		// All of them are reading in a separate goroutine.
		// And will be competing for reads.
		go func(i int, d <-chan int) {
			defer wg.Done()
			for val := range d {
				fmt.Printf("#%d got %d\n", i, val)
			}
		}(i, ch)
	}

	wg.Wait()
}

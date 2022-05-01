package main

import (
	"fmt"
	"sync"
	"time"
)

func Funnel(sources ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	dest := make(chan int)
	wg.Add(len(sources))
	for _, ch := range sources {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				dest <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()
	return dest
}

func main() {
	sources := make([]<-chan int, 0)

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)

		go func() {
			defer close(ch)

			for i := 0; i < 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}

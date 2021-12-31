package main

import "fmt"

func main() {
	channels := [5](chan int){
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go func() {
		// Start to waiting on channels
		for _, ch := range channels {
			fmt.Println("Receiving from", <-ch)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Sending on channel:", i)
		channels[i] <- i
	}
}

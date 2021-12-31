package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan int, 3)
	done := make(chan bool)
	// completed := make(chan bool)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("Received data from channel 1")
			case val := <-ch2:
				fmt.Println(val)
			case c := <-ch3:
				fmt.Println(c)
			case <-done:
				fmt.Println("Exiting...")
				return
			}
		}
	}()

	ch1 <- 100
	ch2 <- "ch2 msg"
	done <- true
}

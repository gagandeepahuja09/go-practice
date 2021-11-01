package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	x++
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	// this is added to give the time to close the channel
	time.Sleep(1 * time.Second)
	val := <-c
	fmt.Println("Read:", val)
	// to give the time to read from the channel
	time.Sleep(1 * time.Second)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

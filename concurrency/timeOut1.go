package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		// timeout won't allow to reach here
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 4):
		// won't reach here as the method will return after 3 seconds
		fmt.Println("Timeout c2")
	}
}

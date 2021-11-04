package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-time.After(4 * time.Second):
			fmt.Printf("\ntime.After()...\n")
		case <-end:
			close(end)
			// question: why is this not printing?
			fmt.Printf("Closing end...")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer value")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers\n", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		// read create number generated and written in the gen goroutine
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	end <- true
}

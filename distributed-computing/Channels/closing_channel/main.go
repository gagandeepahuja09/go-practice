package main

import (
	"fmt"
)

type msg struct {
	id      int
	message string
}

func handleIntChan(intChan <-chan int, done chan<- struct{}) {
	for i := 0; i < 6; i++ {
		fmt.Println(<-intChan)
	}
	done <- struct{}{}
}

func handleMessageChan(msgChan <-chan msg, done chan<- struct{}) {
	for i := 1; i < 7; i++ {
		fmt.Println(<-msgChan)
	}
	done <- struct{}{}
}

func main() {
	intChan := make(chan int)
	done := make(chan struct{})

	go func() {
		intChan <- 5
		intChan <- 7
		intChan <- 9
		close(intChan)
	}()

	go handleIntChan(intChan, done)

	msgChan := make(chan msg, 4)

	go func() {
		for i := 1; i < 6; i++ {
			msgChan <- msg{
				id:      i,
				message: fmt.Sprintf("Message Received %v", i),
			}
		}
		close(msgChan)
	}()

	go handleMessageChan(msgChan, done)

	<-done
	<-done

	intChan <- 100
}

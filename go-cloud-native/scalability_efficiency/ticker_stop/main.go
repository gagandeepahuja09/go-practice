package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(5 * time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-done:
				fmt.Println("Done")
				return
			}
		}
	}()
	<-timer.C
	done <- struct{}{}
}

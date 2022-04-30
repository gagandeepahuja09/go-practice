package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(400 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Println("Ticked at:", time.Now())
			}
		}
	}()
	time.Sleep(1700 * time.Millisecond)
	done <- true
}

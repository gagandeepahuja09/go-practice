package main

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("21:22:03"))
}

func listenToMusic() {
	for {
		printTime("Listening to music...")
	}
}

func completeMail1(wg *sync.WaitGroup) {
	printTime("Completed 1st email")
	wg.Done()
}

func completeMail2(wg *sync.WaitGroup) {
	printTime("Completed 2nd email")
	wg.Done()
}

func completeMail3(wg *sync.WaitGroup) {
	printTime("Completed 3rd email")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go listenToMusic()
	time.Sleep(time.Nanosecond * 4)
	go completeMail1(&wg)
	go completeMail2(&wg)
	go completeMail3(&wg)
	wg.Wait()
}

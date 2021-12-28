package main

import (
	"fmt"
	"sync"
)

// Simple individual tasks
func makeHotelReservation(wg *sync.WaitGroup) {
	fmt.Println("Done making hotel reservation.")
	wg.Done()
}

func bookFlightTickets(wg *sync.WaitGroup) {
	fmt.Println("Done booking flight tickets.")
	wg.Done()
}

func orderADress(wg *sync.WaitGroup) {
	fmt.Println("Done ordering a dress.")
	wg.Done()
}

func payCreditCardBills(wg *sync.WaitGroup) {
	fmt.Println("Done paying credit card bills.")
	wg.Done()
}

// Tasks that will be done in parts
// Only introducing goroutines for the split tasks.

// Writing a mail
func writeAMail(wg *sync.WaitGroup) {
	fmt.Println("Wrote 1/3rd of the mail.")
	go continueWritingMail1(wg)
}

func continueWritingMail1(wg *sync.WaitGroup) {
	fmt.Println("Wrote 2/3rd of the mail.")
	go continueWritingMail2(wg)
}

func continueWritingMail2(wg *sync.WaitGroup) {
	fmt.Println("Done writing the mail.")
	wg.Done()
}

// Listening to audio book
func listenToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Listened first 10 minutes of book")
	go continueListeningToAudioBook(wg)
}

func continueListeningToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Done listening to audio book.")
	wg.Done()
}

var listOfTasks = []func(*sync.WaitGroup){makeHotelReservation, bookFlightTickets, orderADress, payCreditCardBills, writeAMail, listenToAudioBook}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(listOfTasks))
	for _, task := range listOfTasks {
		go task(&waitGroup)
	}
	waitGroup.Wait()
}

package main

import "fmt"

// Simple individual tasks
func makeHotelReservation() {
	fmt.Println("Done making hotel reservation.")
}

func bookFlightTickets() {
	fmt.Println("Done booking flight tickets.")
}

func orderADress() {
	fmt.Println("Done ordering a dress.")
}

func payCreditCardBills() {
	fmt.Println("Done paying credit card bills.")
}

// Tasks that will be done in parts
// Only introducing goroutines for the split tasks.

// Writing a mail
func writeAMail() {
	fmt.Println("Wrote 1/3rd of the mail.")
	go continueWritingMail1()
}

func continueWritingMail1() {
	fmt.Println("Wrote 2/3rd of the mail.")
	go continueWritingMail2()
}

func continueWritingMail2() {
	fmt.Println("Done writing the mail.")
}

// Listening to audio book
func listenToAudioBook() {
	fmt.Println("Listened first 10 minutes of book")
	go continueListeningToAudioBook()
}

func continueListeningToAudioBook() {
	fmt.Println("Done listening to audio book.")
}

var listOfTasks = []func(){makeHotelReservation, bookFlightTickets, orderADress, payCreditCardBills, writeAMail, listenToAudioBook}

func main() {
	for _, task := range listOfTasks {
		task()
	}
}

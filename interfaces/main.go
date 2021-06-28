package main

import "fmt"

type bot interface {
	getGreeting() string
}

func main() {
	sb := spanishBot{}
	eb := englishBot{}
	printGreeting(sb)
	printGreeting(eb)
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

package main

import (
	"fmt"
)

func helloChan(ch <-chan string) {
	val := <-ch
	fmt.Println("Hello ", val)
}

func main() {
	ch := make(chan string)
	go helloChan(ch)
	ch <- "Bob"
}

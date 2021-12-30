package main

import "fmt"

func main() {
	funcChan := make(chan func(int))
	fcn1 := func(i int) {
		fmt.Println("fcn1", i)
	}
	fcn2 := func(i int) {
		fmt.Println("fcn2", i*2)
	}
	fcn3 := func(i int) {
		fmt.Println("fcn3", i*3)
	}
	done := make(chan bool)
	go func() {
		for fcn := range funcChan {
			fcn(10)
		}
		fmt.Println("Exiting")
		done <- true
	}()
	funcChan <- fcn1
	funcChan <- fcn2
	funcChan <- fcn3
	close(funcChan)
	// this is an alternative to using waitgroup, it will keep on waiting to
	// read till something is written on this channel.
	<-done
}

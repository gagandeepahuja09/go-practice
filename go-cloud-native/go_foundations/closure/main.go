package main

import "fmt"

func incrementer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	increment := incrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	increment = incrementer()
	fmt.Println(increment())
}

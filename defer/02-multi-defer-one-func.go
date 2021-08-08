package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		defer fmt.Print(i)
		i++
	}
}

// Since defer uses a stack, hence the largest number at the top
// So, O/P = 54321

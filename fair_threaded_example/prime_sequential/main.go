package main

import (
	"fmt"
	"time"
)

const MAX_INT = 1000000

var primeCount = 0

func checkPrime(x int) {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return
		}
	}
	primeCount++
}

func main() {
	startTime := time.Now()
	for i := 0; i < MAX_INT; i++ {
		checkPrime(i)
	}
	fmt.Println("PROG_TOOK", time.Since(startTime))
	fmt.Println("NUMBER_OF_PRIMES", primeCount)
}

// PROG_TOOK 70.335208ms
// NUMBER_OF_PRIMES 78500

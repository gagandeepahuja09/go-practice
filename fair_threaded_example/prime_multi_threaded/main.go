package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MAX_INT         = 1000000
	MAX_CONCURRENCY = 10
)

var primeCount int32 = 0

func checkPrime(wg *sync.WaitGroup, x int) {
	defer wg.Done()
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&primeCount, 1)
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(MAX_INT)
	for i := 0; i < MAX_INT; i++ {
		go checkPrime(&wg, i)
	}
	wg.Wait()
	fmt.Println("PROG_TOOK", time.Since(startTime))
	fmt.Println("NUMBER_OF_PRIMES", primeCount)
}

// PROG_TOOK 330.725375ms
// NUMBER_OF_PRIMES 78500

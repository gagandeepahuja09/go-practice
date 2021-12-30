package main

import (
	"fmt"
	"sync"
)

func createCashier(cashierID int, wg *sync.WaitGroup) func(int) {
	ordersProcessed := 0
	return func(orderNum int) {
		if ordersProcessed < 3 {
			fmt.Println("Cashier ", cashierID, "Processing order", orderNum, "Orders processed", ordersProcessed)
			ordersProcessed++
		} else {
			fmt.Println("Cashier ", cashierID, "I am tired! I want to take rest!", orderNum)
		}
		wg.Done()
	}
}

func main() {
	cashierIndex := 0
	var wg sync.WaitGroup
	cashiers := []func(int){}
	for i := 1; i <= 3; i++ {
		cashiers = append(cashiers, createCashier(i, &wg))
	}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		cashierIndex = i % 3
		go cashiers[cashierIndex](i)
	}
	wg.Wait()
}

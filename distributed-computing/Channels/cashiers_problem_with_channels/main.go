package main

import (
	"fmt"
	"sync"
)

func cashier(cashierID int, orderChannel <-chan int, wg *sync.WaitGroup) {
	for ordersProcessed := 0; ordersProcessed < 10; ordersProcessed++ {
		orderNum := <-orderChannel
		fmt.Println("Cashier", cashierID, "Processing order", orderNum, "Orders Processed", ordersProcessed)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(30)
	ordersChannel := make(chan int)
	for i := 0; i < 3; i++ {
		go cashier(i, ordersChannel, &wg)
	}
	for i := 0; i < 30; i++ {
		ordersChannel <- i
	}
	wg.Wait()
}

package main

import "log"

func logErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	logErr(err)

	err = vendingMachine.insertMoney(10)
	logErr(err)

	err = vendingMachine.dispenseItem()
	logErr(err)
}

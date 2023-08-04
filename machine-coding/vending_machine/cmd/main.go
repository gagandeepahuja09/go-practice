package main

import (
	"fmt"

	"go-practice.com/machine-coding/vending_machine/vending_machine"
)

func main() {
	// VendingMachine should have all the functions for updating inventory
	// and for carrying out multiple actions
	// v.FuncName should do v.currentState.FuncName
	v := vending_machine.NewVendingMachine()
	fmt.Printf("%+v \n", v)
	fmt.Println(v.ViewInventory())

	// state: 'idle' -> 'no_money' -> 'idle'
	if err := v.PressInsertCashButton(); err != nil {
		fmt.Println("Error Insert first time: ", err)
	}
	if err := v.PressInsertCashButton(); err != nil {
		fmt.Println("Error Insert second time: ", err)
	}
	if coins, err := v.CancelOrRefund(); err != nil {
		fmt.Println("Error Refund first time: ", err)
		fmt.Println("Coins returned: ", coins)
	}

	// state: 'idle' -> 'no_money' -> 'has_money' -> 'idle'
	if err := v.PressInsertCashButton(); err != nil {
		fmt.Println("Error Insert third time: ", err)
	}
	if err := v.InsertCoins(20); err != nil {
		fmt.Println("Error Insert coins first time: ", err)
	}
	if coins, err := v.CancelOrRefund(); err != nil {
		fmt.Println("Error Refund second time: ", err)
		fmt.Println("Coins returned: ", coins)
	}

	// state: 'idle' -> 'no_money' -> 'has_money' -> 'product_selected'
	// -> 'product_dispensed' -> 'idle'
	if err := v.PressInsertCashButton(); err != nil {
		fmt.Println("Error Insert fourth time: ", err)
	}
	if err := v.InsertCoins(20); err != nil {
		fmt.Println("Error Insert coins second time: ", err)
	}
	if err := v.ChooseProductAndQuantity("bro_002", 2); err != nil {
		fmt.Println("Error choose product and quantity first: ", err)
	}
	if err := v.ChooseProductAndQuantity("beverage_001", 6); err != nil {
		fmt.Println("Error choose product and quantity second: ", err)
	}
	if err := v.ChooseProductAndQuantity("beverage_001", 3); err != nil {
		fmt.Println("Error choose product and quantity third: ", err)
	}
	if err := v.ChooseProductAndQuantity("beverage_002", 2); err != nil {
		fmt.Println("Error choose product and quantity fourth: ", err)
	}
	leftCoins, err := v.DispenseProduct()
	fmt.Println("Error dispense produce: ", err)
	fmt.Println("Coins left: ", leftCoins)
	if err := v.PressInsertCashButton(); err != nil {
		fmt.Println("Error Insert fifth time: ", err)
	}
}

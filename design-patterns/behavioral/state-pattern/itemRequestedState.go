package main

import "fmt"

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

func (i *itemRequestedState) insertMoney(money int) error {
	itemPrice := i.vendingMachine.itemPrice
	if money < itemPrice {
		return fmt.Errorf("Please insert amount = %s", itemPrice)
	}
	fmt.Println("Amount inserted is appropriate")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

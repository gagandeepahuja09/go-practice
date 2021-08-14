package main

import "fmt"

type noItemState struct {
	vendingMachine vendingMachine
}

func (i *noItemState) addItem(count) error {
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *noItemState) requestItem() error {
	fmt.Errorf("Item out of stock")
}

func (i *noItemState) dispenseItem() {
	fmt.Errorf("Item out of stock")
}

func (i *noItemState) insertMoney() {
	fmt.Errorf("Item out of stock")
}

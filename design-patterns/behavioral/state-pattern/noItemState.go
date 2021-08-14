package main

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

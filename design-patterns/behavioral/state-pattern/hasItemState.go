package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d Items to be added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("Item must be selected first")
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Item must be selected first")
}

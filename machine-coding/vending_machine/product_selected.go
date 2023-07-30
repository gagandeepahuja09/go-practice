package main

import (
	"go-practice.com/machine-coding/vending_machine/inventory"
)

type productSelected struct {
	vendingMachine *vendingMachine
}

func (s productSelected) PressInsertCashButton() error {
	// s.ctx.SetState(NewproductSelected())
	return nil
}

func (s productSelected) InsertCoins(coins int) error {
	return nil
}

func (s productSelected) ViewInventory() string {
	return inventory.View()
}

func (s productSelected) ChooseProductAndQuantity(productId string, quantity int) error {
	return nil
}

func (s productSelected) DispenseProduct() error {
	return nil
}

func (s productSelected) CancelOrRefund() error {
	return nil
}

func (s productSelected) ReturnChange() error {
	return nil
}

package main

import (
	"go-practice.com/machine-coding/vending_machine/inventory"
)

type productDispensed struct {
	vendingMachine *vendingMachine
}

func (s productDispensed) PressInsertCashButton() error {
	// s.ctx.SetState(NewproductDispensed())
	return nil
}

func (s productDispensed) InsertCoins(coins int) error {
	return nil
}

func (s productDispensed) ViewInventory() string {
	return inventory.View()
}

func (s productDispensed) ChooseProductAndQuantity(productId string, quantity int) error {
	return nil
}

func (s productDispensed) DispenseProduct() error {
	return nil
}

func (s productDispensed) CancelOrRefund() error {
	return nil
}

func (s productDispensed) ReturnChange() error {
	return nil
}

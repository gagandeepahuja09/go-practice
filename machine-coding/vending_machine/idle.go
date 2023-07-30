package main

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type idle struct {
	vendingMachine *vendingMachine
}

func (s idle) PressInsertCashButton() error {
	s.vendingMachine.setState(s.vendingMachine.noMoney)
	return nil
}

func (s idle) InsertCoins(coins int) error {
	return errors.New(ErrOpNotSupported)
}

func (s idle) ViewInventory() string {
	return inventory.View()
}

func (s idle) ChooseProductAndQuantity(productId string, quantity int) error {
	return errors.New(ErrOpNotSupported)
}

func (s idle) DispenseProduct() error {
	return errors.New(ErrOpNotSupported)
}

func (s idle) CancelOrRefund() error {
	return errors.New(ErrOpNotSupported)
}

func (s idle) ReturnChange() error {
	return errors.New(ErrOpNotSupported)
}

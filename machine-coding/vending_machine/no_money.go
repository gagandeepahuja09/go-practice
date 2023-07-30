package main

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type noMoney struct {
	vendingMachine *vendingMachine
}

func (s noMoney) PressInsertCashButton() error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) InsertCoins(coins int) error {
	s.vendingMachine.coins = coins
	s.vendingMachine.setState(s.vendingMachine.hasMoney)
	return nil
}

func (s noMoney) ViewInventory() string {
	return inventory.View()
}

func (s noMoney) ChooseProductAndQuantity(productId string, quantity int) error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) DispenseProduct() error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) CancelOrRefund() error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) ReturnChange() error {
	return errors.New(ErrOpNotSupported)
}

package vending_machine

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type idle struct {
	VendingMachine *VendingMachine
}

func (s idle) PressInsertCashButton() error {
	s.VendingMachine.setState(s.VendingMachine.noMoney)
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

func (s idle) DispenseProduct() (int, error) {
	return 0, errors.New(ErrOpNotSupported)
}

func (s idle) CancelOrRefund() (int, error) {
	return 0, errors.New(ErrOpNotSupported)
}

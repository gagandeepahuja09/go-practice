package vending_machine

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type productSelected struct {
	VendingMachine *VendingMachine
}

func (s productSelected) PressInsertCashButton() error {
	return errors.New(ErrOpNotSupported)
}

func (s productSelected) InsertCoins(coins int) error {
	return errors.New(ErrOpNotSupported)
}

func (s productSelected) ViewInventory() string {
	return inventory.View()
}

func (s productSelected) ChooseProductAndQuantity(productId string, quantity int) error {
	return errors.New(ErrOpNotSupported)
}

func (s productSelected) DispenseProduct() (int, error) {
	changeLeft := s.VendingMachine.coins - s.VendingMachine.productCost*s.VendingMachine.productQuantity
	s.VendingMachine.setState(s.VendingMachine.idle)
	return changeLeft, nil
}

func (s productSelected) CancelOrRefund() (int, error) {
	return s.VendingMachine.hasMoney.CancelOrRefund()
}

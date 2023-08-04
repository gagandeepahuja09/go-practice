package vending_machine

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type noMoney struct {
	VendingMachine *VendingMachine
}

func (s noMoney) PressInsertCashButton() error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) InsertCoins(coins int) error {
	s.VendingMachine.coins = coins
	s.VendingMachine.setState(s.VendingMachine.hasMoney)
	return nil
}

func (s noMoney) ViewInventory() string {
	return inventory.View()
}

func (s noMoney) ChooseProductAndQuantity(productId string, quantity int) error {
	return errors.New(ErrOpNotSupported)
}

func (s noMoney) DispenseProduct() (int, error) {
	return 0, errors.New(ErrOpNotSupported)
}

func (s noMoney) CancelOrRefund() (int, error) {
	return s.VendingMachine.hasMoney.CancelOrRefund()
}

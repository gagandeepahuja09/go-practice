package main

import (
	"errors"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

type hasMoney struct {
	vendingMachine *vendingMachine
}

func (s hasMoney) PressInsertCashButton() error {
	return errors.New(ErrOpNotSupported)
}

func (s hasMoney) InsertCoins(coins int) error {
	return errors.New(ErrOpNotSupported)
}

func (s hasMoney) ViewInventory() string {
	return inventory.View()
}

func (s hasMoney) ChooseProductAndQuantity(productId string, quantity int) error {
	pd := inventory.GetProductDetails(productId)
	// if quantity >
	totalCost := pd.Cost * quantity
}

func (s hasMoney) DispenseProduct() error {
	return nil
}

func (s hasMoney) CancelOrRefund() error {
	return nil
}

func (s hasMoney) ReturnChange() error {
	return nil
}

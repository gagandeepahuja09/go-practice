package vending_machine

import (
	"errors"
	"fmt"

	"go-practice.com/machine-coding/vending_machine/inventory"
)

const (
	ProductNotFoundErr             = "Product with product_id: %s not found"
	ProductQuantityNotAvailableErr = "Product with product_id: %s has max quantity of %d. Please try with lesser quantity"
	InsufficientCoinsErr           = "Insufficient coins to buy product"
)

type hasMoney struct {
	VendingMachine *VendingMachine
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
	pd, exists := inventory.GetProductDetails(productId)
	if !exists {
		return fmt.Errorf(ProductNotFoundErr, productId)
	}
	if quantity > pd.Count {
		return fmt.Errorf(ProductQuantityNotAvailableErr, productId, pd.Count)
	}
	if pd.Cost*quantity > s.VendingMachine.coins {
		return errors.New(InsufficientCoinsErr)
	}
	s.VendingMachine.productCost = pd.Cost
	s.VendingMachine.productQuantity = quantity
	s.VendingMachine.setState(s.VendingMachine.productSelected)
	return nil
}

func (s hasMoney) DispenseProduct() (int, error) {
	return 0, errors.New(ErrOpNotSupported)
}

func (s *hasMoney) resetToIdleState() {
	s.VendingMachine.productCost = 0
	s.VendingMachine.productQuantity = 0
	s.VendingMachine.coins = 0
	s.VendingMachine.setState(s.VendingMachine.idle)
}

func (s hasMoney) CancelOrRefund() (int, error) {
	coinsReturned := s.VendingMachine.coins
	s.resetToIdleState()
	return coinsReturned, nil
}

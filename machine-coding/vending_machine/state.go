package main

const (
	ErrOpNotSupported = "operation not supported in the current state"
)

// todo: this is not the best way to structure code in golang
// try to break in separate packages
type State interface {
	PressInsertCashButton() error
	InsertCoins(coins int) error
	ViewInventory() string
	ChooseProductAndQuantity(productId string, quantity int) error
	DispenseProduct() error
	CancelOrRefund() error
	ReturnChange() error
}

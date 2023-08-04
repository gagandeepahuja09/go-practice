package vending_machine

const (
	ErrOpNotSupported = "operation not supported in the current state"
)

type State interface {
	PressInsertCashButton() error
	InsertCoins(coins int) error
	ViewInventory() string
	ChooseProductAndQuantity(productId string, quantity int) error
	DispenseProduct() (int, error)
	CancelOrRefund() (int, error)
}

package vending_machine

// why can't we create a function: isOperationAllowed which has the state and operation
// mapped. Then you won't even require create implementations for each method of the
// State interface. Can be used when you want to give same response "operation not allowed"
// in all cases

type VendingMachine struct {
	idle             State
	noMoney          State
	hasMoney         State
	productSelected  State
	productDispensed State

	coins           int
	productCost     int
	productQuantity int
	currentState    State
}

func NewVendingMachine() *VendingMachine {
	v := VendingMachine{}
	idle := &idle{
		VendingMachine: &v,
	}
	noMoney := &noMoney{
		VendingMachine: &v,
	}
	hasMoney := &hasMoney{
		VendingMachine: &v,
	}
	productSelected := &productSelected{
		VendingMachine: &v,
	}

	v.idle = idle
	v.noMoney = noMoney
	v.hasMoney = hasMoney
	v.productSelected = productSelected
	v.setState(idle)

	return &v
}

func (v *VendingMachine) PressInsertCashButton() error {
	return v.currentState.PressInsertCashButton()
}

func (v *VendingMachine) InsertCoins(coins int) error {
	return v.currentState.InsertCoins(coins)
}

func (v *VendingMachine) ViewInventory() string {
	return v.currentState.ViewInventory()
}

func (v *VendingMachine) ChooseProductAndQuantity(productId string, quantity int) error {
	return v.currentState.ChooseProductAndQuantity(productId, quantity)
}

func (v *VendingMachine) DispenseProduct() (int, error) {
	return v.currentState.DispenseProduct()
}

func (v *VendingMachine) CancelOrRefund() (int, error) {
	return v.currentState.CancelOrRefund()
}

func (v *VendingMachine) setState(state State) {
	v.currentState = state
}

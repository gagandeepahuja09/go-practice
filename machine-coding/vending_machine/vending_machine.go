package main

// why can't we create a function: isOperationAllowed which has the state and operation
// mapped. Then you won't even require create implementations for each method of the
// State interface. Can be used when you want to give same response "operation not allowed"
// in all cases

type vendingMachine struct {
	idle             State
	noMoney          State
	hasMoney         State
	productSelected  State
	productDispensed State

	coins        int
	currentState State
}

func newVendingMachine() {
	v := vendingMachine{}
	idle := &idle{
		vendingMachine: &v,
	}
	noMoney := &noMoney{
		vendingMachine: &v,
	}
	hasMoney := &hasMoney{
		vendingMachine: &v,
	}
	productSelected := &productSelected{
		vendingMachine: &v,
	}
	productDispensed := &productDispensed{
		vendingMachine: &v,
	}

	v.idle = idle
	v.noMoney = noMoney
	v.hasMoney = hasMoney
	v.productSelected = productSelected
	v.productDispensed = productDispensed
	v.setState(idle)
}

func (v vendingMachine) setState(state State) {
	v.currentState = state
}

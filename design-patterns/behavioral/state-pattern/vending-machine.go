package main

import "fmt"

type vendingMachine struct {
	// we will maintain all the states in the context as well as the current state
	// this current state will keep on changing by the various state implementations
	// by calling vendingMachine's setState method.

	// also the state implementations will keep an instant of the context(vendingMachine) so that they
	// can setState i.e change the state depending on the action(method).

	// Also it is important to remember that the client will directly interact with the vending machine
	// and it need not worry about the current state
	// vendingMachine will do a sort of redirect, eg. vendingMachine.requestItem
	//  ===> vendingMachine.currentState.requestItem
	// now whichever state we are in currently, that state's method will be called.
	hasItem       state
	itemRequested state
	noItem        state
	hasMoney      state

	currentState state

	// for simplicity, let's assume that there will be only one item.
	itemCount int
	itemPrice int
}

// all the states will also store the vending machine instance.
func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}

	// initialize the current state as has item
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.noItem = noItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState

	return v
}

// request item, add item, insert money, dispense item
func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount += count
}

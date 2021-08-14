package main

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) addItem() {
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *itemRequestedState) requestItem() {
}

func (i *itemRequestedState) dispenseItem() {
}

func (i *itemRequestedState) insertMoney() {
}

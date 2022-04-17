package state

import "fmt"

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		_ = fmt.Errorf("inserted money is less. Plesase insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

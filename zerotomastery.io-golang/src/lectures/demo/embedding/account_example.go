package main

import "fmt"

type Account struct {
	accountId int
	balance int
	name string
}

type ManagerAccount struct {
	Account
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprintf("Standard (%v) $%v \"%v\"", 
		a.accountId, 
		a.balance, 
		a.name)
}

func (m ManagerAccount) String() string {
	return fmt.Sprintf("Manager (%v) $%v \"%v\"", 
		m.accountId, 
		m.balance, 
		m.name)
}
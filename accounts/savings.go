package accounts

import (
	"errors"
	"github.com/antunesgabriel/bank_account_manager/clients"
)

type Savings struct {
	Holder clients.Holder
	balance float64
	Operation, Number, Agency int
}

func (c *Savings) Withdraw (amount float64) error {

	if amount > 0 && amount <= c.balance {
		c.balance -= amount
	} else {
		return errors.New("insufficient funds")
	}

	return nil
}

func (c *Savings) Deposit (amount float64) (float64, error) {

	if amount >= 0 {
		c.balance += amount
	} else {
		return c.balance, errors.New("invalid deposit amount")
	}

	return c.balance, nil
}

func (c *Savings) SeeBalance () float64 {
	return c.balance
}
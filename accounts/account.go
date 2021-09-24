package accounts

import (
	"errors"
	"github.com/antunesgabriel/bank_account_manager/clients"
)

type Account struct {
	Holder clients.Holder
	Agency int
	Number int
	balance float64
}

func (c *Account) Withdraw (amount float64) error {

	if amount > 0 && amount <= c.balance {
		c.balance -= amount
	} else {
		return errors.New("insufficient funds")
	}

	return nil
}

func (c *Account) Deposit (amount float64) (float64, error) {

	if amount >= 0 {
		c.balance += amount
	} else {
		return c.balance, errors.New("invalid deposit amount")
	}

	return c.balance, nil
}

func (c *Account) Transfer (amount float64, account *Account) error {
	if amount > 0 && amount <= c.balance {
		c.balance -= amount

		account.balance += amount
	} else {
		return errors.New("it was not possible to carry out the transfer, check the value")
	}

	return nil
}

func (c *Account) SeeBalance () float64 {
	return c.balance
}
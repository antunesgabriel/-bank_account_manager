package accounts

import (
	"errors"
)

type Account struct {
	Holder string
	Agency int
	Number int
	Balance float64
}

func (c *Account) Withdraw (amount float64) error {

	if amount > 0 && amount <= c.Balance {
		c.Balance -= amount
	} else {
		return errors.New("insufficient funds")
	}

	return nil
}

func (c *Account) Deposit (amount float64) (float64, error) {

	if amount >= 0 {
		c.Balance += amount
	} else {
		return c.Balance, errors.New("invalid deposit amount")
	}

	return c.Balance, nil
}

func (c *Account) Transfer (amount float64, account *Account) error {
	if amount > 0 && amount <= c.Balance {
		c.Balance -= amount

		account.Balance += amount
	} else {
		return errors.New("it was not possible to carry out the transfer, check the value")
	}

	return nil
}
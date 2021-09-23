package main

import (
	"fmt"
)

type Account struct {
	holder string
	agency int
	number int
	balance float64
}

func (c *Account) Withdraw (amount float64) string {

	if amount > 0 && amount <= c.balance {
		c.balance -= amount

		return "Withdrawal made!"
	} else {
		return "Insufficient funds!"
	}
}

func main() {
	accountOne := Account{
		holder: "Fantasma",
		agency: 1113,
		number: 4576543,
		balance: 9560.05,
	}

	accountTwo := Account {
		"Nana",
		1113,
		6789422,
		9560.89,
	}

	fmt.Println(accountOne, accountTwo)

	var accountThree *Account

	accountThree = new(Account)

	accountThree.holder = "Gabriel"
	accountThree.agency = 1113
	accountThree.number = 4565652
	accountThree.balance = 0.0

	fmt.Println(accountThree, *accountThree, &accountThree)
	
	fmt.Println(accountOne.Withdraw(560), accountOne)
}
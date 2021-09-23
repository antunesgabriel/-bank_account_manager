package main

import (
	"fmt"
	"github.com/antunesgabriel/bank_account_manager/accounts"
)

func main() {
	accountOne := accounts.Account{
		Holder: "Fantasma",
		Agency: 1113,
		Number: 4576543,
		Balance: 9560.05,
	}

	accountTwo := accounts.Account {
		"Nana",
		1113,
		6789422,
		9560.89,
	}

	fmt.Println(accountOne, accountTwo)

	var accountThree *accounts.Account

	accountThree = new(accounts.Account)

	accountThree.Holder = "Gabriel"
	accountThree.Agency = 1113
	accountThree.Number = 4565652
	accountThree.Balance = 0.0

	fmt.Println(accountThree, *accountThree, &accountThree)
	
	fmt.Println(accountOne.Withdraw(560.05), accountOne)

	balance, err := accountOne.Deposit(1000)

	fmt.Println(balance, err, accountOne) // 10000

	if err = accountOne.Transfer(50000, &accountTwo); err != nil {
		fmt.Println(err)
	}

	fmt.Println(accountOne, accountTwo)

	if err = accountOne.Transfer(1000, &accountTwo);  err != nil {
		fmt.Println(err)
 	}

	fmt.Println(accountOne, accountTwo)
}
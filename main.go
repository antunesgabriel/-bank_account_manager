package main

import (
	"fmt"
	"github.com/antunesgabriel/bank_account_manager/accounts"
	"github.com/antunesgabriel/bank_account_manager/clients"
	"time"
)

type Pay interface {
	Withdraw (amount float64) error
}

func PayBankSlip (account Pay, amount float64) error {
	return account.Withdraw(amount)
}

func main() {
	accountOneHolder := clients.Holder{
		"Nana",
		"0000000091",
		time.Date(2018, time.January, 15, 0, 0, 0, 0, time.UTC),
	}

	accountOne := accounts.Account{
		Holder: accountOneHolder,
		Agency: 1113,
		Number: 4576543,
	}

	accountOne.Deposit(9560.05)

	accountTwoHolder := clients.Holder{
		"Nana",
		"0000000091",
		time.Date(2018, time.January, 15, 0, 0, 0, 0, time.UTC),
	}

	accountTwo := accounts.Account {
		Holder: accountTwoHolder,
		Agency: 1112,
		Number: 6789422,
	}

	accountTwo.Deposit(4000)

	fmt.Println(accountOne, accountTwo)

	var accountThree *accounts.Account

	accountThree = new(accounts.Account)

	accountThree.Holder = clients.Holder{
		"Gabriel",
		"0000000091",
		time.Date(1996, time.December, 22, 0, 0, 0, 0, time.UTC),
	}
	accountThree.Agency = 1113
	accountThree.Number = 4565652
	accountThree.Deposit(0)

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

	 marianaBirth := time.Date(1999, time.July, 3, 0, 0, 0, 0, time.UTC)

	 holderSavings := clients.Holder{
		 "Mariana",
		 "0000000092",
		 marianaBirth,
	 }

	 savingsAccount := accounts.Savings{
		 Holder: holderSavings,
		 Operation: 013,
		 Number: 945945,
	 }

	 savingsAccount.Deposit(450)

	fmt.Println("Saldo na conta 1 e 2", accountOne.SeeBalance(), accountTwo.SeeBalance())

	PayBankSlip(&accountOne, 50)
	 PayBankSlip(&savingsAccount, 400)

	fmt.Println("Saldo da conta 1", accountOne.SeeBalance())

	fmt.Println("Saldo da conta poupança", savingsAccount.SeeBalance())
}
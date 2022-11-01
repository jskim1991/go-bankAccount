package main

import (
	"bankAccount/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("jay")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	withdrawError := account.Withdraw(100)
	if withdrawError != nil {
		fmt.Println(withdrawError)
	} else {
		fmt.Println(account.Balance())
	}

	account.ChangeOwner("John")
	fmt.Println(account)
}

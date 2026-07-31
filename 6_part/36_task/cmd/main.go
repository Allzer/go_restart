package main

import (
	"fmt"
	"go_restart/6_part/36_task/internal/bank"
)

func main() {
	account := bank.BankAccount{
		Owner: "Doni Shapkro",
		Balance: 0,
	}

	err := account.Deposit(-5)
	if err != nil {
		fmt.Println(err)
	}
	err = account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

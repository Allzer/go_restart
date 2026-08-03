package bank

import (
	"go_restart/6_part/36_task/internal/my_errors"
)

type BankAccount struct {
	Owner   string
	Balance int
}

func (b *BankAccount) Deposit(ammount int) error {
	if ammount <= 0 {
		err := myerrors.ErrInvalidAmount
		return err
	}
	b.Balance += ammount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount > b.Balance {
		return myerrors.ErrInsufficientАunds
	}
	b.Balance -= amount
	return nil
}

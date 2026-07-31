package bank

import "errors"

type BankAccount struct {
	Owner   string
	Balance int
}

func (b *BankAccount) Deposit(ammount int) error {
	if ammount <= 0 {
		err := errors.New("Сумма пополнения должна быть больше 0")
		return err
	}
	b.Balance += ammount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount > b.Balance {
		return errors.New("На балансе недостаточно средств")
	}
	b.Balance -= amount
	return nil
}

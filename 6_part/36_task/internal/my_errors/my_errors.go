package myerrors

import "errors"

var ErrInvalidAmount = errors.New("Сумма пополнения должна быть больше 0")
var ErrInsufficientАunds = errors.New("На балансе недостаточно средств")
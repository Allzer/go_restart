package main

import (
	"go_restart/5_part_orders_task/internal/bankcard"
	"go_restart/5_part_orders_task/internal/cryptocurrency"
	"go_restart/5_part_orders_task/internal/order"
	"go_restart/5_part_orders_task/internal/sbp"
)

func main() {

	var sbpPayment order.Payment = sbp.Sbp{}
	var cryptocurrencyPayment order.Payment = cryptocurrency.Cryptocurrency{}
	var bankcardPayment order.Payment = bankcard.Bankcard{}

	orderNumber1 := order.Order{
		ID: 1234,
		TotalPrice: 5678,
		Status: "Ожидает оплаты",
	}

	order.PayForTheOrder(sbpPayment, &orderNumber1)
	order.PayForTheOrder(cryptocurrencyPayment, &orderNumber1)
	order.PayForTheOrder(bankcardPayment, &orderNumber1)
}


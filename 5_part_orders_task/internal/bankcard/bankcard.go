package bankcard

import (
	"fmt"
	"go_restart/5_part_orders_task/internal/order"
)

type Bankcard struct{}

func (b Bankcard) Pay(o *order.Order) bool {
	fmt.Println("Ожидаем подключение к банку")
	return true
}

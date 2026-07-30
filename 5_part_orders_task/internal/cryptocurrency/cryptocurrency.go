package cryptocurrency

import (
	"fmt"
	"go_restart/5_part_orders_task/internal/order"
)

type Cryptocurrency struct{}

func (c Cryptocurrency) Pay(o *order.Order) bool {
	fmt.Println("Ожидаем подключение к кошельку")
	return true
}


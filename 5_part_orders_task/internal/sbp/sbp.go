package sbp

import (
	"fmt"
	"go_restart/5_part_orders_task/internal/order"
)

type Sbp struct{}

func (s Sbp) Pay(o *order.Order) bool {
	fmt.Println("QR-code сгенерирован")
	return true
}

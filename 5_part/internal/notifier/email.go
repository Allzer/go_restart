package notifier

import (
	"fmt"
	order "go_restart/5_part/internal/Order"
)

type EmailNotifier struct {}

func (EmailNotifier) Notify(o *order.Order) {
	fmt.Printf(
		"Чек на Email. Заказ №%d завершён. Сумма заказа: %.2fр",
		o.ID,
		o.TotalPrice,
	)
}
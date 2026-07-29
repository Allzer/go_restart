package notifier

import order "go_restart/5_part/internal/Order"

type Notifier interface{ Notify(*order.Order) }

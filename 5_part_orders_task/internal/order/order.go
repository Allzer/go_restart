package order

type Order struct {
	ID         int
	TotalPrice float64
	Status     string
}

func PayForTheOrder(p Payment, o *Order) {

    if o.Status != "Ожидает оплаты" {
        return
    }

    success := p.Pay(o)

    if success {
        o.Status = "Оплачен"
    }
}


package order

import (
	"go_restart/5_part/internal/Product"
	"math/rand"
)

type Order struct {
	TotalPrice float32
	Status     string
	ID         int
	Products   []product.Product
}

func NewOrder() Order {
	return Order{
		TotalPrice: 0.0,
		Status: "Define",
		ID: rand.Intn(999),
	}
}

func CreateBaseOrder(order Order) Order {
	baseProductList := product.CreateProductList()
	for _, v := range baseProductList {
		order.TotalPrice += v.Price
	}
	order.ID = rand.Intn(999)
	order.Status = "Complited"
	order.Products = baseProductList
	return order
}
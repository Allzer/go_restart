package order

import "go_restart/5_part/internal/Product"

type Order struct {
	TotalPrice float32
	Status     string
	ID         int
	Products   []product.Product
}

func (o *Order) AddProduct(product *product.Product) {
	o.Products = append(o.Products, *product)
}

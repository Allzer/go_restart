package order

import (
	"fmt"
	"go_restart/5_part/internal/Product"
	"math/rand"
)

type Order struct {
	TotalPrice float32
	Status     string
	ID         int
	Products   []product.Product
}

func (o *Order) AddBaseOrder() {
	baseProductList := product.GetProductList()
	for _, v := range baseProductList {
		o.TotalPrice += v.Price
	}
	o.ID = rand.Intn(99999)
	o.Status = "Ожидает подтверждения"
	o.Products = baseProductList
}

func (o *Order)ChooseProduct() {
	baseProductList := product.GetProductList()
	o.ID = rand.Intn(99999)
Menu:
	for {
		param := getMenu()
		switch param {
		case 6:
			o.Status = "Ожидает подтверждения"
			break Menu
		}
		if param > 0 && param < 6 {
			addProductInOrder(o, baseProductList[param-1])
		} else {
			fmt.Println("Такого пункта нет в меню")
		}
	}
}

func addProductInOrder(order *Order, product product.Product) {
	order.Products = append(order.Products, product)
	order.TotalPrice += product.Price
}

func (o *Order) CompleteOrder() {
	o.Status = "Завершён"
}

func getMenu() int {
	var param int
	fmt.Println(
		`
		------------------------------------
		Выберите продукт:
		- 1 Молоко - 110.0р
		- 2 Хлеб - 23.15р
		- 3 Макароны - 45.26р
		- 4 Соль - 17.34р
		- 5 Сахарок - 67.69р
		- 6 Закончить выбор продуктов
		------------------------------------
		`)
	fmt.Scan(&param)
	return param
}

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

func NewOrder() Order {
	return Order{
		TotalPrice: 0.0,
		Status:     "Define",
		ID:         rand.Intn(999),
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

func ChooseProduct(order Order) Order {
	baseProductList := product.CreateProductList()
	order.ID = rand.Intn(99999)
	Menu:
	for {
		param := getMenu()
		switch param {
		case 1:
			order.Products = append(order.Products, baseProductList[0])
			order.TotalPrice += baseProductList[0].Price
		case 2:
			order.Products = append(order.Products, baseProductList[1])
			order.TotalPrice += baseProductList[1].Price
		case 3:
			order.Products = append(order.Products, baseProductList[2])
			order.TotalPrice += baseProductList[2].Price
		case 4:
			order.Products = append(order.Products, baseProductList[3])
			order.TotalPrice += baseProductList[3].Price
		case 5:
			order.Products = append(order.Products, baseProductList[4])
			order.TotalPrice += baseProductList[4].Price
		case 6:
			break Menu
		}
	}
	return order
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

package main

import (
	"fmt"
	"go_restart/5_part/internal/User"
)

func main() {
	user := user.User{
		UserName:   "Doni",
		SurName:    "Shapkro",
		Patronymic: "Serg",
	}

Menu:
	for {
		param := getMenu()
		switch param {
		case 1:
			user.PrintUserInfo()
		case 2:
			user.GetOrder()
		case 3:
			user.Orders.ChooseProduct()
			user.GetOrder()
		case 4:
			user.Orders.AddBaseOrder()
			user.GetOrder()
		case 5:
			user.Orders.CompleteOrder()
			user.GetOrder()
		case 6:
			user.CreateNewOrder()
			user.GetOrder()
		case 7:
			break Menu
		}
	}
}

func getMenu() int {
	var param int
	fmt.Println(
		`
		------------------------------------
		Выберите действие с заказом:
		- 1 Просмотр информации пользователя
		- 2 Просмотреть информацию о заказе
		- 3 Выбрать продукты и добавить их в заказ
		- 4 Добавить Базовый заказ
		- 5 Завершить заказ
		- 6 Очистить заказ
		- 7 Выход
		------------------------------------
		`)
	fmt.Scan(&param)
	return param
}

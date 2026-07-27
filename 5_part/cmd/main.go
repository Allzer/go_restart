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
			user.СhooseProduct()
			user.GetOrder()
		case 3:
			user.AddBaseOrder()
			user.GetOrder()
		case 4:
			user.CompleteOrder()
			user.GetOrder()
		case 5:
			user.GetOrder()
			user.CreateNewOrder()
		case 6:
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
		- 2 Выбрать продукты и добавить их в заказ
		- 3 Добавить Базовый заказ
		- 4 Завершить заказ
		- 5 Создать новый заказ
		- 6 Выход
		------------------------------------
		`)
	fmt.Scan(&param)
	return param
}

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
			user.СhooseProduct()
		case 4:
			user.AddBaseOrder()
		case 5:
			user.CompleteOrder()
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
		- 2 Просмотр списка заказов
		- 3 Выбрать продукты и добавить их в заказ
		- 4 Добавить Базовый заказ
		- 5 Завершить заказ
		- 6 Выход
		------------------------------------
		`)
	fmt.Scan(&param)
	return param
}

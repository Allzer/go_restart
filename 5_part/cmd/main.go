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
	// user.CreateOrder()

	Menu:
	for {
		param := getMenu()
		switch param {
		case 1:
			user.PrintUserInfo()
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var param int
	fmt.Println(
		`
		Выберите действие с заказом:
		- 1 Просмотр информации пользователя
		- 2 Просмотр списка заказов
		- 3 Создать заказ
		- 4 Выход
		`)
	fmt.Scan(&param)
	return param
}

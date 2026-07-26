package user

import (
	"fmt"
	"go_restart/5_part/internal/Order"
)

type User struct {
	UserName   string
	SurName    string
	Patronymic string
	Orders     order.Order
}

func (user *User) PrintUserInfo() {

	if user.UserName != "" && user.SurName != "" && user.Patronymic != "" {
		fmt.Printf("Имя пользователя: %s %s %s", user.UserName, user.SurName, user.Patronymic)
	} else if user.UserName != "" && user.SurName != "" {
		fmt.Printf("Имя пользователя: %s %s", user.UserName, user.SurName)
	} else {
		fmt.Println("Данные пользователя заполнены некорректно")
	}

}

func (user *User) GetOrder() order.Order {
	fmt.Println(user.Orders)
	return user.Orders
}

func (user *User) CreateOrder() {
	user.Orders = order.NewOrder()
}

func (user *User) AddBaseOrder() {
	user.Orders = order.CreateBaseOrder(user.Orders)
}

func (user *User) chooseProduct() {
	
}

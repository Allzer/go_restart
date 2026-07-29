package user

import (
	"fmt"
	"go_restart/5_part/internal/Order"
)

type User struct {
	UserName   string
	SurName    string
	Patronymic string
	Order      order.Order
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

func (user *User) GetOrder() {
	fmt.Printf("Заказ номер %d на сумму %.2f получил статус %s", user.Order.ID, user.Order.TotalPrice, user.Order.Status)
}

func (user *User) CreateNewOrder() {
	user.Order = order.Order{}
}

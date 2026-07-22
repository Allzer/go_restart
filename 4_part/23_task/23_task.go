package main

import "fmt"

type User struct {
	Name    string
	Surname string
}

func main() {
	user := User{
		Name:    "Ltybc",
		Surname: "Ifgrfhby",
	}
	fmt.Println(user)
	SetSystemUser(&user)

	fmt.Println(user)
}

func SetSystemUser(user *User) {
	user.Name = "System"
	user.Surname = "System"
}

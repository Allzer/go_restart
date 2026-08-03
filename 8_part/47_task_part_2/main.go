package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

var ErrInvalidUserName = errors.New("Имя пользователя не должно быть пустым")

func main() {
	var wg sync.WaitGroup
	users := []User{
		{ID: 1, Name: ""},
		{ID: 2, Name: "Alex"},
		{ID: 3, Name: "Maria"},
		{ID: 4, Name: "John"},
	}

	wg.Add(len(users))

	for _, v := range users {
		go func(user User) {
			defer wg.Done()
			err := ValidateUser(user)
			if err == nil {
				ProcessUser(user)
			}
		}(v)
	}
	wg.Wait()
}

func ProcessUser(user User) {
	fmt.Printf("Пользователь %s начал обрабатываться\n", user.Name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Пользователь %s закончил обрабатываться\n", user.Name)
}

func ValidateUser(user User) error {
	if len(user.Name) <= 1 {
		fmt.Printf("Пользователь c ID%d не прошёл проверку имени", user.ID)
		fmt.Println(ErrInvalidUserName)
		return ErrInvalidUserName
	}
	return nil
}

package main

import (
	"fmt"
	"time"
)

type User struct {
	ID  int
	FIO string
}

type Cache struct {
	Data map[int]User
}

var cache = Cache{Data: map[int]User{}}

func main() {
	fmt.Println(getUser(1234))
	fmt.Println(getUser(1234))
}

func cachingInMemory(u User) {
	cache.Data[u.ID] = u
}

func getUser(userId int) User {
	// Эмуляция получения данных из БД
	if user, ok := cache.Data[userId]; ok {
		fmt.Println("Достаём из кеша")
		return user
	}
	fmt.Println("Ожидайте, делаем запрос к бд")
	time.Sleep(10 * time.Second)
	user := newUser(userId)
	cachingInMemory(user)
	return user
}

func newUser(userId int) User {
	return User{
		ID:  userId,
		FIO: "Denis Shapkarin Sergeevich",
	}
}

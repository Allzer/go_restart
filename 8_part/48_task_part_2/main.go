package main

import "fmt"

func main() {
	ch := make(chan string)
	go SendMessages(ch)

	for message := range ch {
		fmt.Println(message)
	}
	fmt.Println("Все сообщения получены")
}

func SendMessages(ch chan string) {
	messages := []string{
		"Привет",
		"Как дела?",
		"Go очень классный!",
		"Пока!",
	}
	for _, v := range messages {
		ch <- v
	}
	close(ch)
}

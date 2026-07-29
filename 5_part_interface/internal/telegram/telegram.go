package telegram

import "fmt"

type TelegramNotify struct{}

func (TelegramNotify) Notify(message string) {
	fmt.Println("[Telegram] ", message)
}
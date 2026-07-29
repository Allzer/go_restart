package main

import (
	"go_restart/5_part_interface/internal/console"
	"go_restart/5_part_interface/internal/email"
	"go_restart/5_part_interface/internal/notifier"
	"go_restart/5_part_interface/internal/sms"
	"go_restart/5_part_interface/internal/telegram"
)

func main() {
	notifiers := []notifier.Notifier{
		console.ConsoleNotify{},
		email.EmailNotify{},
		sms.SmsNotify{},
		telegram.TelegramNotify{},
	}

	for _, v := range notifiers {
		notifier.SendNotification(v, "ваш заказ готов")
	}
}

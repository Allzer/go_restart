package sms

import "fmt"

type SmsNotify struct{}

func (SmsNotify) Notify(message string) {
	fmt.Println("[SMS] ", message)
}

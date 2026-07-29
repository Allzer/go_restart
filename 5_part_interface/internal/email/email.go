package email

import "fmt"

type EmailNotify struct{}

func (EmailNotify) Notify(message string) {
	fmt.Println("[Email] ", message)
}

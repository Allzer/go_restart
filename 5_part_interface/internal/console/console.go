package console

import "fmt"

type ConsoleNotify struct{}

func (ConsoleNotify) Notify(message string) {
	fmt.Println("[Console] ", message)
}

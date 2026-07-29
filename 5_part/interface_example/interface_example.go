package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct{}
func (Cat) Speak() {
	fmt.Println("Miay")
}

type Dog struct{}
func (Dog) Speak() {
	fmt.Println("Gaf")
}

func Speak(s Speaker) {
	s.Speak()
}

func main() {
	Speak(Cat{})
	Speak(Dog{})
}

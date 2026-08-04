package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	ch := make(chan int)

	for _, v := range numbers {
		go Square(v, ch)
	}

	for i := 0; i < len(numbers); i++ {
		result := <-ch
		fmt.Println(result)
	}
}

func Square(number int, ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- number * number
}
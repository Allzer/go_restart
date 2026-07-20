package main

import "fmt"

func main() {
	sumToN(6)
}

func sumToN(n int) {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum = sum + i
		fmt.Println(sum)
	}
}

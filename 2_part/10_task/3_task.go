package main

import "fmt"

func main() {
	factorialN(5)
}

func factorialN(n int) {
	var x int = 1
	for i := 1; i <= n; i++ {
		x = x * i
	}
	fmt.Println(x)
}
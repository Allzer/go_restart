package main

import "fmt"

func main() {
	result1 := calulation(5, 2, sum)
	fmt.Println(result1)
	result2 := calulation(5, 2, multiply)
	fmt.Println(result2)
	result3 := calulation(2, 5, minus)
	fmt.Println(result3)
}

func calulation(a, b int, operation func(a, b int) int) int {
	return operation(a, b)
}

func sum(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }
func minus(a, b int) int    { return a - b }

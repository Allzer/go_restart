package main

import "fmt"

func main() {
	sum := sumPrams(2, 2)
	fmt.Println(sum)
}

func sumPrams(a, b int) int {
	return a + b
}

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 76, 7}
	changeArrElements(arr)
	fmt.Println(arr)
}

func changeArrElements(arr []int) {
	for i := range arr {
		(arr)[i] *= 2
	}
}

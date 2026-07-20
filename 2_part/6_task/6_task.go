package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 10, 12, 12, 14}
	popDuble(arr)
}

func popDuble(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	fmt.Println(arr)
}

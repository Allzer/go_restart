package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := genArr()
}

func genArr() []int {
	arr := make([]int, 10)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	return arr
}

func popDuble(arr []int) {
	for i:= 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1]{
			fmt.Println(arr[i])
		}
	}
}

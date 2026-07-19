package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := genArr()
	reversArr(arr)
}

func genArr() []int {
	arr := make([]int, 10)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	return arr
}

func reversArr(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
}

package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	arr := genArr()
	arr_2 := genArr()

	slices.Sort(arr_2)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
	fmt.Println(arr_2)

}

func genArr() []int {
	arr := make([]int, 10)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	return arr
}
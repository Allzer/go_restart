package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := genArr()
	minMaxAvg(arr)
}

func genArr() []int {
	arr := make([]int, 10)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	return arr
}

func minMaxAvg(arr []int) {

	var min int
	var max int
	var avg int 

	for _, v := range arr {
		min = v
		if v < min{
			min = v
		} else if v > max {
			max = v
		}
		avg += v
	}
	avg = avg/len(arr)
	fmt.Println(min, max, avg)
}


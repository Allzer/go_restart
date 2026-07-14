package main

import "fmt"

func main() {
	max := FindMax([]int{1, -5, -3, -4, -3, -6})
	fmt.Println(max)
}

func FindMax(arr []int) int {
	var max int = arr[0]

	if len(arr) == 0 {
		return 0
	}
	
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

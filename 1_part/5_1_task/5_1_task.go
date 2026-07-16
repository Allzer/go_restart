package main

import "fmt"

func main() {
	var arr = []int{4, 5, 321, 76, 5, 4, 123, 4, 5, 6}

	bybbleSort(arr)
	fmt.Println(arr)
}

func bybbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

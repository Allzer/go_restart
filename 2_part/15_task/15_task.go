package main

import "fmt"

func main() {
	elementCounter([]int{1, 2, 3, 3, 3, 2, 2, 3, 3})
}

func elementCounter(arr []int) {
	counter := make(map[int]int)

	for _, v := range arr {
		counter[v]++
	}

	fmt.Println(counter)
}

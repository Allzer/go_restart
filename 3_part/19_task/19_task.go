package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 6, 7}
	fmt.Println(arr)
	operationWithArr(mapArr, arr)
	operationWithArr(filterArr, arr)
	operationWithArr(reduceArr, arr)
}

func operationWithArr(operation func([]int), arr []int) {
	operation(arr)
}

func mapArr(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for i := range newArr {
		newArr[i] = newArr[i] * 2
	}
	fmt.Println(newArr)
}

func filterArr(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for i := 0; i < len(newArr)-1; i++ {
		if newArr[i]%2 == 0 {
			newArr = append(newArr[:i], newArr[i+1:]...)
			i-- //для того, чтобы не првыгать через индекс удалённого элемента
		}
	}
	fmt.Println(newArr)
}

func reduceArr(arr []int) {
	var totalSum int
	
	for _, v := range arr {
		totalSum += v
	}
	fmt.Println(totalSum)
}

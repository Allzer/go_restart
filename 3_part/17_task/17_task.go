package main

import "fmt"

func main() {
	sum := oneReturn(2, 3)
	fmt.Println(sum)

	stringa := multiReturn("X D")
	fmt.Println(stringa)

	allSum1 := variadicPrams(1, 2, 3, 4)
	allSum2 := variadicPrams(1, 3, 4)
	allSum3 := variadicPrams([]int{1, 2, 3, 4, 6}...)
	fmt.Println(allSum1)
	fmt.Println(allSum2)
	fmt.Println(allSum3)

}

func oneReturn(a, b int) int { return a + b }

func multiReturn(a string) string {
	if a == "XD" {
		return "lmao"
	}
	return "not LMAo"
}

func variadicPrams(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

package main

import "fmt"

func main() {
	s := make([]int, 3, 5)

	for i := range s {
		s[i] = i + 1
	}

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	
	s = append(s, 4, 5, 6, 7, 8, 98, 0)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

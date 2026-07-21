package main

import (
	"fmt"
	"restart/3_part/calc"
)

func main() {
	a := calc.Sum(2, 3)
	b := calc.Minus(2, 3)
	c := calc.Multiply(2, 3)
	d := calc.Division(2, 3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

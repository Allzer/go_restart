package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	
	changeX(&x)
	fmt.Println(x)
}

func changeX(x *int) {
	*x = 20
}

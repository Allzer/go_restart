package rectangle

import "fmt"

type Rectangle struct {
	A int
	B int
}

func (r Rectangle) Area(...int) {
	fmt.Println("Площадь прямоугольника равна ", r.A*r.B)
}

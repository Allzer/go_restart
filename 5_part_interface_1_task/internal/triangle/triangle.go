package triangle

import (
	"math"
)

type Triangle struct {
	A int
	B int
	C int
}

func (t Triangle) Area() float32 {
	p := (t.A+t.B+t.C)/2
	area := math.Sqrt(float64((p*(p-t.A)*(p-t.B)*(p-t.C))))
	return float32(area)
}
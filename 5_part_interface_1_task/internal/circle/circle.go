package circle

import (
	"math"
)

type Circle struct {
	R int
}

func (c Circle) Area() float32 {
	return math.Pi * float32(c.R*c.R)
}

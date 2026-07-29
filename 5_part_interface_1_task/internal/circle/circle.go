package circle

import (
	"fmt"
	"math"
)

type Circle struct {
	R int
}

func (c Circle) Area(...int) {
	fmt.Println("Площадь окружности равна ", math.Pi*float64(c.R*c.R))
}
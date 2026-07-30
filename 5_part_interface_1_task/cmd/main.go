package main

import (
	"fmt"
	"go_restart/5_part_interface_1_task/internal/circle"
	"go_restart/5_part_interface_1_task/internal/rectangle"
	"go_restart/5_part_interface_1_task/internal/shape"
	"go_restart/5_part_interface_1_task/internal/triangle"
)

func main() {
	figures := []shape.Shape{
		rectangle.Rectangle{A: 4, B: 2},
		circle.Circle{R: 6},
		triangle.Triangle{A: 3, B: 4, C: 5},
	}

	for _, v := range figures {
		fmt.Printf("Площадь равна %.2f \n", v.Area())
	}
}
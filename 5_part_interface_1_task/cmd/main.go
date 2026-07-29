package main

import (
	"go_restart/5_part_interface_1_task/internal/circle"
	"go_restart/5_part_interface_1_task/internal/rectangle"
	"go_restart/5_part_interface_1_task/internal/shapeInterface"
	"go_restart/5_part_interface_1_task/internal/triangle"
)

func main() {
	figures := []shapeInterface.Shape{
		rectangle.Rectangle{4,2},
		circle.Circle{6},
		triangle.Triangle{3,4,5},
	}

	for _, v := range figures{
		v.Area()
	}
}
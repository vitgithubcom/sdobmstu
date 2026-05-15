package main

import (
	"shapes/internal/entity"
	"shapes/internal/service"
)

func main() {

	shapes := []entity.Shape{
		entity.Circle{Radius: 5.0},
		entity.Rectangle{Width: 4.0, Height: 6.0},
		entity.Circle{Radius: 3.5},
		entity.Rectangle{Width: 7.0, Height: 2.0},
	}
	
	service.PrintAreas(shapes)
}
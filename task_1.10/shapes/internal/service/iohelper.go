package service

import (
	"fmt"
	"shapes/internal/entity"
)

func PrintAreas(shapes []entity.Shape) {
	fmt.Println("=== Площади фигур ===\n")
	
	for i, shape := range shapes {
		area := shape.Area()
		
		switch s := shape.(type) {
		case entity.Circle:
			fmt.Printf("%d. Круг (радиус = %.2f): площадь = %.2f\n", i+1, s.Radius, area)
		case entity.Rectangle:
			fmt.Printf("%d. Прямоугольник (ширина = %.2f, высота = %.2f): площадь = %.2f\n", 
				i+1, s.Width, s.Height, area)
		default:
			fmt.Printf("%d. Фигура: площадь = %.2f\n", i+1, area)
		}
	}
}
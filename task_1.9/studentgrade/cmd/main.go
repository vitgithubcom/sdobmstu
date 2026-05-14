package main

import (
	"fmt"
	"studentgrade/internal/logic"
)

func main() {
	student := logic.Student{
		Name:   "Иван Петров",
		Grades: []float64{4, 5, 3, 2, 5},
	}
	
	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Оценки: %v\n", student.Grades)
	fmt.Printf("Средний балл: %.2f\n", student.AverageGrade())
}
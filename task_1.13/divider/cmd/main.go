package main

import (
	"divider/internal/logic"
	"fmt"
)

func main() {
	examples := []struct {
		a, b float64
	}{
		{10, 2},
		{5, 0},
		{7.5, 2.5},
		{-10, 2},
	}

	for _, ex := range examples {
		result, err := logic.Divide(ex.a, ex.b)
		if err != nil {
			fmt.Printf("%.2f / %.2f = Ошибка: %v\n", ex.a, ex.b, err)
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", ex.a, ex.b, result)
		}
	}
}
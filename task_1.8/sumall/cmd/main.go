package main

import (
	"fmt"
	"sumall/internal/logic"
)

func main() {
	fmt.Printf("sumAll(1, 2, 3) = %d\n", logic.SumAll(1, 2, 3))
	fmt.Printf("sumAll(10, -2, 4, 7) = %d\n", logic.SumAll(10, -2, 4, 7))
	fmt.Printf("sumAll(5, 5, 5, 5, 5) = %d\n", logic.SumAll(5, 5, 5, 5, 5))
	fmt.Printf("sumAll() = %d\n", logic.SumAll())
}
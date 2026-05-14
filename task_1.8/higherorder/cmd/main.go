package main

import (
	"fmt"
	"higherorder/internal/logic"
)

func main() {
	a, b := 10, 5
	
	sum := logic.ApplyOperation(a, b, logic.Add)
	diff := logic.ApplyOperation(a, b, logic.Subtract)
	prod := logic.ApplyOperation(a, b, logic.Multiply)
	
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d - %d = %d\n", a, b, diff)
	fmt.Printf("%d * %d = %d\n", a, b, prod)
}
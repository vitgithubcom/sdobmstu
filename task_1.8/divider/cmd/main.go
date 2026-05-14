package main

import (
	"divider/internal/logic"
	"divider/internal/service"
)

func main() {
	a, b := service.ReadNumbers()
	
	result, err := logic.Divide(a, b)
	
	service.PrintResult(a, b, result, err)
}
package main

import (
	"fmt"
	"closure/internal/logic"
)

func main() {
	counter := logic.CreateCounter()
	
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
	fmt.Println(counter()) // 4
}
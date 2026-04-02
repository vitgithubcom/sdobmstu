package main

import (
	"fmt"
	"math"
)

func main() {
	b := 0.1
	b += 0.2
	
	fmt.Println(math.Abs(b-0.3)<0.0001)
}
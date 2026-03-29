package main

import "fmt"

func main() {	

	var a int
	a = 3
	switch {
	case a > 1:
		fmt.Println("a > 1")		
	case a == 2:
		fmt.Println("2")
	case a == 3 || a == 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}	

}
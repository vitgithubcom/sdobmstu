package main

import "fmt"

func factorial(n uint64) uint64 {
	result := uint64(1)
	for i := uint64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var num uint64
	fmt.Scan(&num)
	fmt.Printf("Факториал числа %d = %d\n", num, factorial(num))
}
package service

import (
	"fmt"
)

func PrintResults(arr [10]int, sortedSlice []int) {
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Отсортированный слайс:", sortedSlice)
}
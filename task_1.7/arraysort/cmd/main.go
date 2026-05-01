package main

import (
	"arraysort/internal/logic"
	"arraysort/internal/service"
)

func main() {
	
	arr := logic.GenerateRandomArray(10, 1, 100)
	
	sortedSlice := logic.CopyAndSort(arr)
	
	service.PrintResults(arr, sortedSlice)
}
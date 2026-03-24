package logic

import (
	"sort"
)

// SortDescending сортирует числа по убыванию
func SortDescending(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

// FindMax находит максимальное число
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// FindMin находит минимальное число
func FindMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

// CalculateAverage вычисляет среднее арифметическое
func CalculateAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}
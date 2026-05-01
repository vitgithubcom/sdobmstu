package logic

import (
	"math/rand"
	"sort"
)

func GenerateRandomArray(size, min, max int) ([10]int) {
	
	var arr [10]int
	
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}
	
	return arr
}

func CopyAndSort(arr [10]int) []int {
	slice := make([]int, len(arr))
	copy(slice, arr[:])
	
	sort.Ints(slice)
	
	return slice
}

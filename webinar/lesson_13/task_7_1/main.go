package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	str := ""
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			mu.Lock()
			str += strconv.Itoa(index)
			mu.Unlock()
			wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println(str)
}
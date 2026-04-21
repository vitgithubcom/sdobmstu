package main

import (
    "fmt"
    "sync"    
)


var (
	mu   sync.Mutex
	next int = 1
)
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mu.Lock()
		if id == next {
			fmt.Println(id)
			next++
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var currentTurn int = 0
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				mu.Lock()
				if currentTurn == id {
					fmt.Println(id)
					currentTurn++
					mu.Unlock()
					break
				}
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

}
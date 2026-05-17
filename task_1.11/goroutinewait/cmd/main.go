package main

import (
	"fmt"
	"goroutinewait/internal/logic"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go logic.Worker(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("\nВсе горутины завершили работу. Программа завершена.")
}
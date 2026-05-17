package logic

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("Горутина %d начала работу\n", id)
	
	time.Sleep(time.Millisecond * 500)
	
	fmt.Printf("Горутина %d завершила работу\n", id)
}
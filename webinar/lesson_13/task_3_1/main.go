package main

import (
	"fmt"
	"sync"
)

func channelReader(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go channelReader(ch, &wg)
	
	for i := 0; i < 99; i++ {
		ch <- i
	}
	
	wg.Wait()
}
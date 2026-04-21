package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func sleepyGopher(index int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Duration(rand.IntN(5) * int(time.Second)))
	fmt.Println(index)
	
}

func main() {
	
	var wg sync.WaitGroup
	//wg.Add(5)

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go sleepyGopher(i, &wg)
	}

	wg.Wait()

}	
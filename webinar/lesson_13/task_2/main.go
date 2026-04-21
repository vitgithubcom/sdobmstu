package main

import (
	"math/rand/v2"
	"fmt"
	"time"
)

func sleepyGopher(index int, ch chan int){
	time.Sleep(time.Duration(rand.IntN(6) * int(time.Second)))
	ch <- index
}

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, ch)
	}

	timeout := time.After(3 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case id := <-ch:			
			fmt.Printf("Gopher #%d finished\n", id)	
		case <-timeout:
			fmt.Println("Timeout, sorry =()")	
			return		
		}
	}


}	
package main

import (
	"fmt"	
	"time"
)

func channelReader(ch chan int){
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	go channelReader(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	
	
	time.Sleep(100*time.Second)
}	
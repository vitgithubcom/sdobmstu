package main

import (
	"math/rand/v2"
	"fmt"
	"time"
)

func sleepyGopher(index int){
	time.Sleep(time.Duration(rand.IntN(10) * int(time.Second)))
	fmt.Println(index)
}

func main() {
	
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}

	time.Sleep(10*time.Second)

}	
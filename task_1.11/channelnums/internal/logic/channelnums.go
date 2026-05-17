package logic

import (
	"fmt"
)

func Generator(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(ch chan int) {
	fmt.Println("Полученные числа:")
	
	for num := range ch {
		fmt.Printf("%d ", num)
	}
}
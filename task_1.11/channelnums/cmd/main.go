package main

import (
	"channelnums/internal/logic"
)

func main() {
	
	ch := make(chan int)
	
	go logic.Generator(ch)
	
	logic.Consumer(ch)
}
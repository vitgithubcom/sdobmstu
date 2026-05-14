package main

import (
	"fmt"
	"bookinfo/internal/logic"
)

func main() {
	
	book := logic.Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}
	
	
	fmt.Println(book.GetInfo())
}
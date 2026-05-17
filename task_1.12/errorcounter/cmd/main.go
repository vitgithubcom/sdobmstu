package main

import (
	"errorcounter/internal/logic"
	"fmt"
)

func main() {
	filename := "data/logfile.log.txt"
	
	count := logic.CountErrorsInFile(filename)
	
	fmt.Printf("Количество строк со словом 'error': %d\n", count)
}
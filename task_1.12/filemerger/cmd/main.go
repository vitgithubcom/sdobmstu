package main

import (
	"filemerger/internal/logic"
	"fmt"
)

func main() {
	files := logic.GetTextFiles("data")
	
	if len(files) == 0 {
		fmt.Println("Текстовые файлы не найдены в папке data/")
		return
	}
	
	fmt.Println("Найдены файлы:", files)
	
	logic.MergeFiles(files, "data", "combined.txt")
	
	fmt.Println("Готово! Результат в data/combined.txt")
}
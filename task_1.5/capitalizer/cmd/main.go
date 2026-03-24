package main

import (
	"log"
	"capitalizer/internal/logic"
	"capitalizer/internal/service"
	"os"
)

func main() {
	// 1. Получаем строку от пользователя
	input, err := service.ReadString()
	if err != nil {
		log.Printf("Ошибка ввода: %v\n", err)
		os.Exit(1)
	}
	
	// 2. Преобразуем слова
	result := logic.CapitalizeWords(input)
	
	// 3. Выводим результат
	service.PrintResult(input, result)
}
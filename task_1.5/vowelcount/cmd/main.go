package main

import (
	"log"
	"vowelcount/internal/logic"
	"vowelcount/internal/service"
	"os"
)

func main() {
	// 1. Получаем строку от пользователя
	input, err := service.ReadString()
	if err != nil {
		log.Printf("Ошибка ввода: %v\n", err)
		os.Exit(1)
	}
	
	// 2. Подсчитываем количество гласных
	count := logic.CountVowels(input)
	
	// 3. Выводим результат
	service.PrintResult(input, count)
}
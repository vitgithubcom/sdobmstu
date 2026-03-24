package main

import (
	"log"
	"strings-counter/internal/service"
	"os"
)

func main() {
	// 1. Получаем строку от пользователя
	input, err := service.ReadString()
	if err != nil {
		log.Printf("Ошибка ввода: %v\n", err)
		os.Exit(1)
	}
	
	// 2. Подсчитываем количество символов
	count := len(input)
	
	// 3. Выводим результат
	service.PrintResult(input, count)
}
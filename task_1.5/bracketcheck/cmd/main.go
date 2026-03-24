package main

import (
	"log"
	"bracketcheck/internal/logic"
	"bracketcheck/internal/service"
	"os"
)

func main() {
	// 1. Получаем строку от пользователя
	input, err := service.ReadString()
	if err != nil {
		log.Printf("Ошибка ввода: %v\n", err)
		os.Exit(1)
	}
	
	// 2. Проверяем правильность расстановки скобок
	openCount, closeCount, isValid := logic.CheckBrackets(input)
	
	// 3. Выводим результат
	service.PrintResult(input, openCount, closeCount, isValid)
}
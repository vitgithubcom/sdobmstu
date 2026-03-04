package main

import (
	"fmt"
	"log"
	"numbers-sorter/internal/logic"
	"numbers-sorter/internal/service"
	"os"
)

func main() {
	// 1. Получаем числа от пользователя
	numbers, err := service.ReadNumbers()
	if err != nil {
		log.Printf("Ошибка ввода: %v\n", err)
		os.Exit(1)
	}
	
	// 2. Сортируем по убыванию
	sorted := logic.SortDescending(numbers)
	
	// 3. Вычисляем параметры
	max := logic.FindMax(numbers)
	min := logic.FindMin(numbers)
	avg := logic.CalculateAverage(numbers)
	
	// 4. Выводим результаты
	service.PrintResults(sorted, max, min, avg)
}
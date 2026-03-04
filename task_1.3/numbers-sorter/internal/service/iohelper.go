package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadNumbers читает 5 натуральных чисел из стандартного ввода
func ReadNumbers() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите 5 натуральных чисел через пробел: ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ввода: %v", err)
	}
	
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	
	if len(parts) != 5 {
		return nil, fmt.Errorf("требуется ровно 5 чисел, получено %d", len(parts))
	}
	
	numbers := make([]int, 0, 5)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования числа %d: %v", i+1, err)
		}
		
		if num <= 0 {
			return nil, fmt.Errorf("число %d должно быть натуральным (> 0)", num)
		}
		
		numbers = append(numbers, num)
	}
	
	return numbers, nil
}

// PrintResults выводит результаты на экран
func PrintResults(sorted []int, max, min int, avg float64) {
	fmt.Print("Отсортированные элементы: ")
	for i, num := range sorted {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
	
	fmt.Printf("Самое большое число: %d\n", max)
	fmt.Printf("Самое маленькое число: %d\n", min)
	fmt.Printf("Среднее арифметическое: %.0f\n", avg)
}
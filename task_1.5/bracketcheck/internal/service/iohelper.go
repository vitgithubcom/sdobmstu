package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadString читает строку из стандартного ввода
func ReadString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку-формулу: ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ввода: %v", err)
	}
	
	input = strings.TrimRight(input, "\r\n")
	
	return input, nil
}

// PrintResult выводит результат на экран
func PrintResult(input string, openCount, closeCount int, isValid bool) {
	if isValid {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", openCount, closeCount)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся\n", openCount, closeCount)
	}
}
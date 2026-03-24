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
	fmt.Print("Введите строку: ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ввода: %v", err)
	}
	
	input = strings.TrimRight(input, "\r\n")
	
	return input, nil
}

// PrintResult выводит результат на экран
func PrintResult(input string, count int) {
	fmt.Printf("Количество гласных букв в строке \"%s\": %d\n", input, count)
}
package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumbers() (float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Введите первое число (делимое): ")
	input1, _ := reader.ReadString('\n')
	
	fmt.Print("Введите второе число (делитель): ")
	input2 , _ := reader.ReadString('\n')
	
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)
	
	a, _ := strconv.ParseFloat(input1, 64)
	
	b, _ := strconv.ParseFloat(input2, 64)
	
	return a, b
}

func PrintResult(a, b, result float64, err error) {
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	}
}
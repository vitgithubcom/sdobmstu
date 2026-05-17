package logic

import (
	"bufio"
	"os"
	"strings"
)

func CountErrorsInFile(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	count := 0
	
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "error") {
			count++
		}
	}
	
	return count
}
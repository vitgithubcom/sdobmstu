package service

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadText() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	return input
}


func PrintStatistics(stats map[string]int, totalWords int) {
	fmt.Println("\n=== Статистика частоты слов ===")
	fmt.Printf("Всего слов: %d\n", totalWords)
	fmt.Printf("Уникальных слов: %d\n", len(stats))
	fmt.Println("\nЧастота вхождения каждого слова:")
	
	words := make([]string, 0, len(stats))
	for word := range stats {
		words = append(words, word)
	}
	sort.Strings(words)
	
	for _, word := range words {
		count := stats[word]
		fmt.Printf("  %s: %d раз(а)\n", word, count)
	}
	
	maxWord := ""
	maxCount := 0
	for word, count := range stats {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}
	
	if maxWord != "" {
		fmt.Printf("\nСамое часто встречающееся слово: \"%s\" (%d раз(а))\n", maxWord, maxCount)
	}
}
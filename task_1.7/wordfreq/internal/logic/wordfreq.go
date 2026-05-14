package logic

import (
	"strings"
	"unicode"
)

func CountWordFrequencies(text string) map[string]int {
	wordCount := make(map[string]int)
	
	words := strings.Fields(text)
	
	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		if word != "" {
			wordCount[word]++
		}
	}
	
	return wordCount
}

func GetTotalWordsCount(stats map[string]int) int {
	total := 0
	for _, count := range stats {
		total += count
	}
	return total
}

func GetUniqueWordsCount(stats map[string]int) int {
	return len(stats)
}
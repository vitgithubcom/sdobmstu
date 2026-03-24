package logic

import (
	"strings"
	"unicode"
)

// CapitalizeWords преобразует каждое слово так, чтобы первая буква была заглавной
func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	
	for i, word := range words {
		if len(word) > 0 {
			// Преобразуем первую букву в заглавную, остальные в строчные
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	
	return strings.Join(words, " ")
}
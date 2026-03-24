package logic

import (
	"strings"
)

// CountVowels подсчитывает количество гласных букв в строке
func CountVowels(s string) int {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
	count := 0
	
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	
	return count
}
package main

import (
	"wordfreq/internal/logic"
	"wordfreq/internal/service"
)

func main() {
	text := service.ReadText()
	
	wordStats := logic.CountWordFrequencies(text)
	
	service.PrintStatistics(wordStats, logic.GetTotalWordsCount(wordStats))
} 
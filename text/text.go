package text

import (
	"strings"
)

func ToWords(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int, len(words))
	for _, word := range words {
		wordCount[word]++
	} 
	return wordCount
}

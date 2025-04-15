package codedaily

import (
	"fmt"
)

func mostFrequentWord(words []string) (string, int) {
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	maxCount := 0
	var mostFrequent string
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			mostFrequent = word
		}
	}

	return mostFrequent, maxCount
}

func Main1() {
	words := []string{"go", "java", "go", "python", "go", "java"}
	word, count := mostFrequentWord(words)
	fmt.Printf("%s appears %d times\n", word, count)
}

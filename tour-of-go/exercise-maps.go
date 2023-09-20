// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.

package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	wordsCount := make(map[string]int)
	words := strings.FieldsFunc(s, unicode.IsSpace)
	for _, w := range words {
		wordsCount[w]++
	}
	return wordsCount
}

func RunMapsExercice() {
	wc.Test(wordCount)
	fmt.Println()
}

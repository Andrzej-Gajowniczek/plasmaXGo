package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "syllabus"
	syllables := divideIntoSyllables(word)
	fmt.Println("Word:", word)
	fmt.Println("Syllables:", syllables)
}

func divideIntoSyllables(word string) []string {
	vowels := "aeiouAEIOU"
	syllables := make([]string, 0)

	start := 0
	for i, char := range word {
		if strings.ContainsRune(vowels, char) {
			if i > 0 && !strings.ContainsRune(vowels, rune(word[i-1])) {
				syllables = append(syllables, word[start:i])
				start = i
			}
		}
	}
	syllables = append(syllables, word[start:])

	return syllables
}

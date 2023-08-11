package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed "other/message.txt"
var list string

func breakDowm4syllabes(s string) []string {
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

func main() {

	//text := "Hello, this is a sample text."
	words := strings.Fields(list)
	width := 8
	sheetLineByLine := []string
	lineIndex := 0
	line := ""
	for i, word := range words {

		if len(world) > width {
			syllabes := breakDowm4syllabes(word)
			for num, syllab := range syllabes {

			}

		}

		fmt.Printf("% 3d:%s\n", i, word)

	}

}

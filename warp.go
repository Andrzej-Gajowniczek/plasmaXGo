package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac nisl et arcu sollicitudin elementum."
	maxWidth := 8

	wrappedText := wrapTextToWords(text, maxWidth)
	fmt.Println(wrappedText)
}

func wrapTextToWords(text string, maxWidth int) string {
	words := strings.Fields(text)
	lines := make([]string, 0)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxWidth {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	fmt.Println(lines)
	return strings.Join(lines, "\n")
}

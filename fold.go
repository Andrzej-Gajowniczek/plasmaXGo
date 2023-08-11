package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed "message.txt"
var info string

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac nisl et arcu sollicitudin elementum."
	maxWidth := 40

	wrappedText := wrapAndCenterText(string, maxWidth)
	fmt.Println(wrappedText)
}

func wrapAndCenterText(text string, maxWidth int) string {
	var result string

	words := strings.Fields(text)
	line := ""
	for _, word := range words {
		if len(line)+len(word)+1 > maxWidth {
			padding := (maxWidth - len(line)) / 2
			result += strings.Repeat(" ", padding) + line + "\n"
			line = ""
		}
		if line != "" {
			line += " "
		}
		line += word
	}

	padding := (maxWidth - len(line)) / 2
	result += strings.Repeat(" ", padding) + line

	return result
}

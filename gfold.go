package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed "message.txt"
var text string

func main() {
	//text := "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit.\nSed ac nisl et arcu sollicitudin elementum."
	maxWidth := 10

	wrappedText := wrapAndCenterText(text, maxWidth)
	fmt.Println(wrappedText)
}

func wrapAndCenterText(text string, maxWidth int) string {
	var result string

	paragraphs := strings.Split(text, "\n")
	for _, paragraph := range paragraphs {
		words := strings.Fields(paragraph)
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
		result += strings.Repeat(" ", padding) + line + "\n"
	}

	return result
}

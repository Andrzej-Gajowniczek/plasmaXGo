package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed "other/message.txt"
var info string

func main() {
	maxWidth := 8

	wrappedText := wrapAndCenterText(info, maxWidth)
	fmt.Println(wrappedText)
}

func wrapAndCenterText(text string, maxWidth int) string {
	var result string

	words := strings.Fields(text)
	if len(words) == 0 {
		return result
	}
	//fmt.Println(words)
	os.Exit(1)
	line := words[0]
	for i := 1; i < len(words); i++ {
		if len(line)+len(words[i])+1 > maxWidth {
			padding := (maxWidth - len(line)) / 2
			if padding > 0 {
				result += strings.Repeat(" ", padding) + line + "\n"
			} else {
				result += line + "\n"
			}
			line = words[i]
		} else {
			line += " " + words[i]
		}
	}

	padding := (maxWidth - len(line)) / 2
	if padding > 0 {
		result += strings.Repeat(" ", padding) + line
	} else {
		result += line
	}

	return result
}

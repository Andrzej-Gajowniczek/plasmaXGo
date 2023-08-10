package main

import (
	"fmt"
	"strconv"
)

func main() {
	asciiChars := []string{
		"\u2588", // Full block
		"\u2593", // Dark shade
		"\u2592", // Medium shade
		"\u2591", // Light shade
	}
	hexString := "1A" // Replace this with your hexadecimal string

	// Convert hexadecimal string to uint
	number, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hexadecimal string: %s\n", hexString)
	fmt.Printf("Uint value: %d\n", number)
	fmt.Println(asciiChars)
}

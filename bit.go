package main

import (
	"fmt"
)

func main() {
	originalByte := byte(0b11010101) // Example input byte
	bitSlice := byteToBitSlice(originalByte, 8)

	fmt.Printf("Original Byte: %08b\n", originalByte)
	fmt.Printf("Bit Slice:    %d\n", bitSlice)
}

func byteToBitSlice(b, p byte) []byte {
	bitSlice := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		bit := (b >> uint(i)) & 1
		bitSlice[7-i] = byte(bit * p)
	}
	return bitSlice
}

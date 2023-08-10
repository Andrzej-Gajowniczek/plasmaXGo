package main

import (
	"fmt"
)

func main() {
	/*	err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
	*/
	tablica := make([]uint8, 256, 256)
	var a, b uint8 = 0x40, 255
	fmt.Printf("a:%d; b:%d\n", a, b)
	a--
	b = b + 6
	c := a * b

	fmt.Printf("a:%d; b:%d\n c:=%d type:%T\n", a, b, c, c)
	tablica[0xff] = 111
	fmt.Println(tablica)

}

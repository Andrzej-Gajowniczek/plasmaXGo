package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {

	termbox.Init()
	x, y := termbox.Size()
	termbox.Close()
	fmt.Println("x:", x, "y:", y)
}

package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	// Initialize Termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Your location slice (example)
	var location []termbox.Cell

	// Replace this with your actual location data
	// location = ...

	// Get the terminal dimensions
	width, height := termbox.Size()

	// Create the screen buffer
	screenBuffer := make([]termbox.Cell, width*height)

	// Copy the location data to the screen buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			if index < len(location) {
				screenBuffer[index] = location[index]
			}
		}
	}

	// Clear the screen and draw the buffer
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	termbox.SetCell(x, y, ch, fg, bg)

	// Flush the screen to display the content
	termbox.Flush()

	// Wait for a key press or an event (optional)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Handle key press events here
			// For example, you can exit the loop on pressing 'q'
			if ev.Ch == 'q' {
				return
			}
		case termbox.EventError:
			// Handle errors if any
			panic(ev.Err)
		}
	}
}

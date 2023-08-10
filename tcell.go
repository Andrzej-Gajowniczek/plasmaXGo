package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

func main() {
	// Initialize the Tcell screen
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	// Create a view
	view := views.NewBoxLayout(views.Vertical)

	// Create a textview and add it to the view
	textView := views.NewTextView()
	textView.SetText("Hello, World!")
	view.AddWidget(textView, 0.5)

	// Set the view as the root
	root := views.NewBoxLayout(views.Horizontal)
	root.AddWidget(view, 0.5)

	// Set the root layout as the root of the screen
	screen.SetRoot(root, true)

	// Start the event loop
	for {
		screen.Show()
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

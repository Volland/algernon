package term

import (
	"github.com/gdamore/tcell/termbox"
)

// Default color/attribute settings
var (
	BG          = Blue
	TITLECOLOR  = Cyan | Bold
	TEXTCOLOR   = Black
	BOXBG       = White
	BOXLIGHT    = White | Bold
	BOXDARK     = Black
	BUTTONFOCUS = Yellow | Bold
	BUTTONTEXT  = White | Bold
	LISTFOCUS   = Red
	LISTTEXT    = Black
)

// Place text at the given x and y coordinate.
// fg is the foreground color, while bg is the background color.
func Write(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for pos, r := range text {
		termbox.SetCell(x+pos, y, r, fg, bg)
	}
}

// Place text at the given x and y coordinate,
// using the default color scheme.
func Say(x int, y int, text string) {
	for pos, r := range text {
		termbox.SetCell(x+pos, y, r, TEXTCOLOR, BG)
	}
}

// Remove all text. Clear the screen.
func Clear() {
	termbox.Clear(TEXTCOLOR, BG)
}

// Retrieve the number of character columns available on the current screen
func ScreenWidth() int {
	return First(termbox.Size)
}

// Retrieve the number of lines of characters available on the current screen
func ScreenHeight() int {
	return Second(termbox.Size)
}

// Update the screen with what has been written so far
func Flush() {
	termbox.Flush()
}

// Initialize the text screen (using curses)
func Init() error {
	return termbox.Init()
}

// Close the text screen (using curses)
func Close() {
	termbox.Close()
}

// Retrieve the next event in the queue (like a keypress)
func PollEvent() *termbox.Event {
	e := termbox.PollEvent()
	return &e
}

// Wait for Esc, Enter or Space to be pressed
func WaitForKey() {
	for {
		e := PollEvent()
		switch e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyEsc, termbox.KeyEnter, termbox.KeySpace:
				return
			}
		}
	}
}

// Set the text/forground color
func SetFg(fg termbox.Attribute) {
	TEXTCOLOR = fg
}

// Set the background color
func SetBg(bg termbox.Attribute) {
	BG = bg
}

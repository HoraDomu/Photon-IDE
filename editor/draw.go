package editor

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Default background
var defaultBg = termbox.ColorBlue

// Draw renders the editor with syntax highlighting
func (e *Editor) Draw() {
	termbox.Clear(termbox.ColorWhite, defaultBg)
	width, height := termbox.Size()

	// Ensure at least one line exists
	if len(e.Rows) == 0 {
		e.Rows = append(e.Rows, "")
	}

	// Clamp cursor within visible window
	if e.Cy < e.RowOffset {
		e.RowOffset = e.Cy
	}
	if e.Cy >= e.RowOffset+height-1 {
		e.RowOffset = e.Cy - height + 2
	}

	if e.Cy >= len(e.Rows) {
		e.Cy = len(e.Rows) - 1
	}
	if e.Cy < 0 {
		e.Cy = 0
	}
	if e.Cx > len(e.Rows[e.Cy]) {
		e.Cx = len(e.Rows[e.Cy])
	}
	if e.Cx < 0 {
		e.Cx = 0
	}

	// Draw each visible line
	for y := 0; y < height-1; y++ {
		rowIdx := e.RowOffset + y
		if rowIdx >= len(e.Rows) {
			continue
		}
		line := e.Rows[rowIdx]
		gutter := fmt.Sprintf("%4d ", rowIdx+1)
		drawString(0, y, gutter, termbox.ColorLightGreen, defaultBg)

		x := len(gutter)
		inString := false

		for i := 0; i < len(line); i++ {
			r := rune(line[i])
			fg := termbox.ColorBlack

			// Comments
			if i < len(line)-1 && line[i] == '/' && line[i+1] == '/' {
				for j := i; j < len(line); j++ {
					termbox.SetCell(x, y, rune(line[j]), termbox.ColorRed, defaultBg)
					x++
				}
				break
			}

			// Strings
			if r == '"' {
				inString = !inString
				fg = termbox.ColorYellow
			}
			if inString {
				fg = termbox.ColorYellow
			}

			// Keywords
			if !inString && ((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
				wordStart := i
				for i < len(line) && ((line[i] >= 'a' && line[i] <= 'z') || (line[i] >= 'A' && line[i] <= 'Z')) {
					i++
				}
				word := line[wordStart:i]
				kwColor := KeywordColor(word)
				for _, r2 := range word {
					termbox.SetCell(x, y, r2, kwColor, defaultBg)
					x++
				}
				i--
				continue
			}

			// Normal character
			termbox.SetCell(x, y, r, fg, defaultBg)
			x++
		}
	}

	// Status bar
	status := e.Status
	if e.Dirty {
		status += " [+]"
	}
	drawString(0, height-1, padRight(status, width), termbox.ColorWhite, defaultBg)

	// Set cursor
	termbox.SetCursor(e.Cx+5, e.Cy-e.RowOffset)
	termbox.Flush()
}

// Draws a string with foreground and background
func drawString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

// Pads string with spaces to match width
func padRight(s string, width int) string {
	for len(s) < width {
		s += " "
	}
	if len(s) > width {
		return s[:width]
	}
	return s
}

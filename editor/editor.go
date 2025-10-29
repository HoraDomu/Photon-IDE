package editor

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Editor struct {
	Rows      []string
	Filename  string
	Cx, Cy    int
	RowOffset int
	ColOffset int
	Dirty     bool
	Status    string
	StatusAt  time.Time
}

func NewEditor() *Editor {
	return &Editor{
		Rows: []string{""},
	}
}

func (e *Editor) SetStatus(s string) {
	e.Status = s
	e.StatusAt = time.Now()
}

// Editor main loop
func (e *Editor) Run() {
	for {
		e.Draw()
		e.HandleInput()
	}
}

// Basic editing methods
func (e *Editor) insertRune(r rune) {
	row := e.Rows[e.Cy]
	row = row[:e.Cx] + string(r) + row[e.Cx:]
	e.Rows[e.Cy] = row
	e.Cx++
	e.Dirty = true
}

func (e *Editor) insertNewline() {
	row := e.Rows[e.Cy]
	newRow := row[e.Cx:]
	e.Rows[e.Cy] = row[:e.Cx]
	e.Rows = append(e.Rows[:e.Cy+1], append([]string{newRow}, e.Rows[e.Cy+1:]...)...)
	e.Cy++
	e.Cx = 0
	e.Dirty = true
}

func (e *Editor) backspace() {
	if e.Cx == 0 {
		if e.Cy == 0 {
			return
		}
		prevLen := len(e.Rows[e.Cy-1])
		e.Rows[e.Cy-1] += e.Rows[e.Cy]
		e.Rows = append(e.Rows[:e.Cy], e.Rows[e.Cy+1:]...)
		e.Cy--
		e.Cx = prevLen
	} else {
		row := e.Rows[e.Cy]
		e.Rows[e.Cy] = row[:e.Cx-1] + row[e.Cx:]
		e.Cx--
	}
	e.Dirty = true
}

// Prompt input for filename
func (e *Editor) PromptInput() string {
	var input string
	for {
		e.SetStatus("Enter filename: " + input)
		e.Draw()

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEnter:
				return input
			case termbox.KeyEsc:
				return ""
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			default:
				if ev.Ch != 0 {
					input += string(ev.Ch)
				}
			}
		}
	}
}

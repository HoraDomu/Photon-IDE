package editor

import "github.com/nsf/termbox-go"

func (e *Editor) HandleInput() {
	ev := termbox.PollEvent()
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyArrowLeft:
			if e.Cx > 0 {
				e.Cx--
			}
		case termbox.KeyArrowRight:
			if e.Cx < len(e.Rows[e.Cy]) {
				e.Cx++
			}
		case termbox.KeyArrowUp:
			if e.Cy > 0 {
				e.Cy--
			}
		case termbox.KeyArrowDown:
			if e.Cy+1 < len(e.Rows) {
				e.Cy++
			}
		case termbox.KeyCtrlS:
			if e.Filename == "" {
				filename := e.PromptInput()
				if filename != "" {
					e.Filename = filename
				} else {
					e.SetStatus("Save cancelled")
					break
				}
			}
			if err := e.Save(); err != nil {
				e.SetStatus("Error saving file: " + err.Error())
			} else {
				e.SetStatus("File saved!")
				e.Dirty = false
			}
		case termbox.KeyCtrlQ:
			termbox.Close()
			panic("Quit")
		case termbox.KeySpace:
			e.insertRune(' ')
		case termbox.KeyTab:
			for i := 0; i < 4; i++ {
				e.insertRune(' ')
			}
		case termbox.KeyEnter:
			e.insertNewline()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			e.backspace()
		default:
			if ev.Ch != 0 {
				e.insertRune(ev.Ch)
			}
		}
	}
}

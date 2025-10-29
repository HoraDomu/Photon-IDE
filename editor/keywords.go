package editor

import "github.com/nsf/termbox-go"

func KeywordColor(word string) termbox.Attribute {
	switch word {
	case
		"package", "import", "func", "var", "const",
		"type", "struct", "interface", "return",
		"if", "else", "for", "range", "switch", "case",
		"default", "break", "continue", "go", "defer",
		"map", "chan", "select", "fallthrough":
		return termbox.ColorCyan
	case "true", "false", "nil":
		return termbox.ColorMagenta
	default:
		return termbox.ColorWhite
	}
}

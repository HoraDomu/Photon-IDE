package main

import (
	"fmt"
	"log"
	"os"

	"Photon_v0.1/editor"
	"github.com/nsf/termbox-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./photon <filename>")
		return
	}

	filename := os.Args[1]

	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	ed, err := editor.OpenFile(filename)
	if err != nil {
		ed = editor.NewEditor()
		ed.Filename = filename
	}

	ed.SetStatus("Editing: " + filename)
	ed.Run()
}

package editor

import (
    "bufio"
    "os"
)

func (e *Editor) OpenFile(filename string) {
    f, err := os.Open(filename)
    if err != nil {
        e.Rows = []string{""}
        e.SetStatus("New file")
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    e.Rows = []string{}
    for scanner.Scan() {
        e.Rows = append(e.Rows, scanner.Text())
    }
    if len(e.Rows) == 0 {
        e.Rows = []string{""}
    }
    e.Filename = filename
    e.SetStatus("Opened " + filename)
}

func (e *Editor) SaveFile() {
    if e.Filename == "" {
        e.SetStatus("No filename")
        return
    }
    f, err := os.Create(e.Filename)
    if err != nil {
        e.SetStatus("Save error")
        return
    }
    defer f.Close()

    for _, row := range e.Rows {
        f.WriteString(row + "\n")
    }
    e.Dirty = false
    e.SetStatus("Saved " + e.Filename)
}


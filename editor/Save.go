package editor

import (
	"os"
)

func (e *Editor) Save() error {
	if e.Filename == "" {
		e.Filename = "untitled.txt" //or we can prompt for a filename
	}
	file, err := os.Create(e.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range e.Rows {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

package editor

import (
	"bufio"
	"os"
)

func OpenFile(path string) (*Editor, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	e := &Editor{
		Rows:     []string{},
		Filename: path,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e.Rows = append(e.Rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Ensure at least one empty line exists
	if len(e.Rows) == 0 {
		e.Rows = append(e.Rows, "")
	}

	return e, nil
}

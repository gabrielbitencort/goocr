package modules

import "os"

const filename = "output.txt"

func SaveText(text string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			return
		}
	}(file)

	if _, err := file.WriteString(text + "\n"); err != nil {
		return err
	}

	return nil
}

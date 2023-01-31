package utils

import (
	"bufio"
	"os"
)

func ReplaceFirstLineInFile(file, line, newline string) error {
	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	var bytes []byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if str == line {
			str = newline
		}
		bytes = append(bytes, []byte(str)...)
		bytes = append(bytes, []byte("\n")...)
	}
	f.Close()

	os.WriteFile(file, bytes, 0666)

	return nil
}

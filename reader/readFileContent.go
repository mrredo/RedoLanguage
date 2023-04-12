package reader

import (
	"errors"
	"io"
	"os"
	"strings"
)

func ReadFileContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()
	splitted := strings.Split(filename, ".")
	if splitted[len(splitted)-1] != "rd" {
		return "", errors.New("invalid file extension")
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return "", errors.New("Error reading file:" + err.Error())
	}

	// Print the contents of the file
	return string(content), nil
}

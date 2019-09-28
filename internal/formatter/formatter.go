package formatter

import (
	"errors"
	"io/ioutil"
	"strings"
	"fmt"
)

func Format(file string, width int) (string, error) {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading file (%s): %s", file, err))
	}
	text := strings.Split(string(fileContents), "\n")
	return strings.Join(text, "\n"), nil
}

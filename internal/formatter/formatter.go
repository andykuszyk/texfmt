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
	formattedText := []string{}
	for _, line := range text {
		for _, splitLine := range splitLines(line, width) {
			formattedText = append(formattedText, splitLine)
		}
	}
	return strings.Join(formattedText, "\n"), nil
}

func splitLines(line string, width int) []string {
	if len(line) <= width {
		return []string{line}
	}
	splitIndex := findSplitIndex(line, width)
	splitA := line[:splitIndex]
	splitB := line[splitIndex:]
	if len(splitB) <= width {
		return []string{splitA, splitB}
	} else {
		return append([]string{splitA}, splitLines(splitB, width)...)
	}
}

func findSplitIndex(line string, width int) int {
	index := width
	modifier := 0
	for string(line[index]) != " " {
		index -= 1
		modifier = 1
	}
	return index + modifier
}

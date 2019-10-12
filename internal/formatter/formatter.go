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
	lines := strings.Split(string(fileContents), "\n")
	formattedLines := formatLines(lines, width)
	for {
		shuffled, shuffledLines := shuffleLines(formattedLines, width)
		if !shuffled {
			break
		}
		formattedLines = make([]string, len(shuffledLines))
		copy(formattedLines, shuffledLines)
	}
	return strings.Join(removeEmpty(formattedLines), "\n"), nil
}

func removeEmpty(lines []string) []string {
	nonEmptyLines := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nonEmptyLines = append(nonEmptyLines, line)
	}
	return nonEmptyLines
}

func shuffleLines(lines []string, width int) (bool, []string) {
	shuffledLines := make([]string, len(lines))
	copy(shuffledLines, lines)
	shuffled := false
	for lineIndex, line := range shuffledLines {
		if lineIndex == len(shuffledLines) - 1 {
			break
		}

		if len(line) == width {
			continue
		}

		nextLine := lines[lineIndex + 1]
		if len(nextLine) == 0 {
			continue
		}
		nextWord := getFirstWord(nextLine)
		if len(line) + len(nextWord) <= width {
			shuffledLines[lineIndex] = fmt.Sprintf("%s%s", strings.ReplaceAll(line, "\n", ""), nextWord)
			shuffledLines[lineIndex + 1] = strings.TrimLeft(nextLine, nextWord)
			shuffled = true
			break
		}
	}
	return shuffled, shuffledLines
}

func getFirstWord(line string) string {
	startingIndex := 0
	endingIndex := 1
	for {
		if endingIndex >= len(line) {
			break
		}
		if string(line[endingIndex - 1]) == " " {
			if endingIndex == 1 {
				break
			} else {
				endingIndex -= 1
				break
			}
		}
		endingIndex += 1
	}
	return line[startingIndex:endingIndex]
}

func formatLines(lines []string, width int) []string {
	formattedLines := []string{}
	for _, line := range lines {
		for _, splitLine := range splitLines(line, width) {
			formattedLines = append(formattedLines, splitLine)
		}
	}
	return formattedLines
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

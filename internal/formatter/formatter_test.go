package formatter

import (
	"testing"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"path/filepath"
)

type TestCase struct {
	Name		string
	Before		string
	After		string
}

func TestFormat(t *testing.T) {
	fileInfos, err := ioutil.ReadDir("../../testdata")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []TestCase{}
	for _, beforeFile := range getBeforeFiles(fileInfos) {
		name := getName(beforeFile.Name())
		afterFile := getAfterFile(name, fileInfos)
		if afterFile == nil {
			t.Log(fmt.Sprintf("Skipping test case %s, because a corresponding after file could not be found to the before file", name))
			continue
		}
		testCase := TestCase {
			Name:	name,
			Before: filepath.Join("..", "..", "testdata", beforeFile.Name()),
			After:  getText(afterFile),
		}
		testCases = append(testCases, testCase)
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual, err := Format(testCase.Before, 120)
			require.Nil(t, err)
			assert.Equal(t, testCase.After, actual)
		})
	}
}

func getText(fileInfo os.FileInfo) string {
	bytes, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", fileInfo.Name()))
	if err != nil {
		return ""
	}
	return string(bytes)
}

func getName(fileName string) string {
	index := strings.Index(fileName, "-")
	return fileName[:index]
}

func getAfterFile(name string, fileInfos []os.FileInfo) os.FileInfo {
	for _, fileInfo := range fileInfos {
		if strings.Contains(fileInfo.Name(), name) {
			return fileInfo
		}
	}
	return nil
}

func getBeforeFiles(fileInfos []os.FileInfo) []os.FileInfo {
	beforeFiles := []os.FileInfo{}
	for _, fileInfo := range fileInfos {
		if strings.Contains(fileInfo.Name(), "before") {
			beforeFiles = append(beforeFiles, fileInfo)
		}
	}
	return beforeFiles
}

type FindSplitIndexTestCase struct {
	Name		string
	Line		string
	Index		int
	SplitChar	string
}

func TestFindSplitIndex(t *testing.T) {
	testCases := []FindSplitIndexTestCase{
		FindSplitIndexTestCase{
			Name: "Space after index",
			Line: "aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute",
			Index: 120,
			SplitChar: " ",
		},
		FindSplitIndexTestCase{
			Name: "First space before width",
			Line: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
			Index: 116,
			SplitChar: "a",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := findSplitIndex(testCase.Line, 120)
			assert.Equal(t, testCase.Index, actual)
			assert.Equal(t, testCase.SplitChar, string(testCase.Line[actual]))
		})
	}
}

func TestGetFirstWord(t *testing.T) {
	actual := getFirstWord(" foo bar")
	assert.Equal(t, " ", actual)

	actual = getFirstWord("foo bar")
	assert.Equal(t, "foo", actual)
}

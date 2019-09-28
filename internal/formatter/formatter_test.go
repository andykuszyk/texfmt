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
		actual, err := Format(testCase.Before, 120)
		require.Nil(t, err)
		assert.Equal(t, testCase.After, actual)
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

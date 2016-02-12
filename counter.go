package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CountLines(importPath string) (int, error) {
	rootDir := filepath.Clean(filepath.Join(os.Getenv("GOPATH"), "src", importPath))

	lineCount := 0

	err := filepath.Walk(rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || !strings.HasSuffix(path, ".go") {
				return nil
			}

			linesInFile, err := countLinesInFile(path)
			lineCount += linesInFile
			return err
		})

	return lineCount, err
}

func countLinesInFile(filepath string) (int, error) {
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 0, err
	}
	return bytes.Count(fileContents, []byte("\n")), nil
}

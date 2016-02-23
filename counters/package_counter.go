package counters

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type directoryLister interface {
	ListFiles(dirPath string) ([]string, error)
}

//go:generate counterfeiter -o ../mocks/file_lines_counter.go --fake-name FileLinesCounter . fileLinesCounter
type fileLinesCounter interface {
	CountLines(filePath string) (int, error)
}

type PackageLinesCounter struct {
	GoPath           string
	DirectoryLister  directoryLister
	FileLinesCounter fileLinesCounter
}

func checkPathIsClean(path string) error {
	if strings.Contains(path, "..") || filepath.IsAbs(path) {
		return errors.New("unclean path")
	}
	return nil
}

func (c *PackageLinesCounter) Count(packagePath string) (int, error) {
	if err := checkPathIsClean(packagePath); err != nil {
		return -1, err
	}

	absPackagePath := filepath.Join(c.GoPath, "src", packagePath)
	files, err := c.DirectoryLister.ListFiles(absPackagePath)
	if err != nil {
		return -1, fmt.Errorf("listing files in package: %s", err)
	}

	total := 0
	for _, file := range files {
		lines, err := c.FileLinesCounter.CountLines(file)
		if err != nil {
			return -1, fmt.Errorf("counting lines in %q: %s", file, err)
		}
		total += lines
	}

	return total, nil
}

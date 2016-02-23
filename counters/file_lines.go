package counters

import (
	"bytes"
	"io/ioutil"
)

type FileLinesCounter struct{}

func (c *FileLinesCounter) CountLines(filePath string) (int, error) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return -1, err
	}
	return bytes.Count(fileContents, []byte("\n")), nil
}

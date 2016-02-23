package counters

import (
	"os"
	"path/filepath"
	"strings"
)

type DirectoryLister struct{}

func (l *DirectoryLister) ListFiles(pkgPath string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(pkgPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || !strings.HasSuffix(path, ".go") {
				return nil
			}

			files = append(files, path)
			return err
		})

	if err != nil {
		return nil, err
	}

	return files, nil
}

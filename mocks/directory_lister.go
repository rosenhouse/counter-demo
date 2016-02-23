package mocks

type DirectoryLister struct {
	ListFilesCall struct {
		Receives struct {
			DirPath string
		}
		Returns struct {
			Files []string
			Error error
		}
	}
}

func (l *DirectoryLister) ListFiles(dirPath string) ([]string, error) {
	l.ListFilesCall.Receives.DirPath = dirPath
	return l.ListFilesCall.Returns.Files, l.ListFilesCall.Returns.Error
}

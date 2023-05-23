package pathutils

import (
	"os"
)

type PathType int

const (
	Incorrect PathType = iota
	File
	Folder
)

func CheckPathType(filepath string) PathType {
	fileInfo, statErr := os.Stat(filepath)
	if statErr != nil {
		return Incorrect
	}

	if fileInfo.IsDir() {
		return Folder
	}

	return File
}

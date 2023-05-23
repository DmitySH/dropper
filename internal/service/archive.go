package service

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	SenderErr = errors.New("error on sending side")
)

type ArchiveService struct {
	archiveName string
}

func NewArchiveService(archiveName string) *ArchiveService {
	return &ArchiveService{archiveName: archiveName}
}

func (a *ArchiveService) FolderToTempZIPArchive(folderPath string) (string, error) {
	tmpDirPath, createTempDirErr := os.MkdirTemp("", "dropper")
	if createTempDirErr != nil {
		log.Println("can't create temp directory:", createTempDirErr)
		return "", SenderErr
	}

	archiveFile, createArchiveFileErr := os.Create(filepath.Join(tmpDirPath, a.archiveName))
	if createArchiveFileErr != nil {
		log.Println("can't create archive:", createArchiveFileErr)
		_ = os.RemoveAll(tmpDirPath)
		return "", SenderErr
	}
	defer archiveFile.Close()

	zw := zip.NewWriter(archiveFile)
	defer zw.Close()

	walker := func(path string, info os.FileInfo, recErr error) error {
		if recErr != nil {
			return recErr
		}
		if info.IsDir() {
			return nil
		}

		file, openErr := os.Open(path)
		if openErr != nil {
			return openErr
		}
		defer file.Close()

		zipPath, _ := strings.CutPrefix(path, folderPath)
		zipPath = strings.Trim(zipPath, "\\./")

		zipFile, createZipErr := zw.Create(zipPath)
		if createZipErr != nil {
			return createZipErr
		}

		_, fileToZipErr := io.Copy(zipFile, file)
		if fileToZipErr != nil {
			return fileToZipErr
		}

		return nil
	}

	walkErr := filepath.Walk(folderPath, walker)
	if walkErr != nil {
		log.Println("error during recursive archiving:", walkErr)
		_ = os.RemoveAll(tmpDirPath)
		return "", walkErr
	}

	return tmpDirPath, nil
}

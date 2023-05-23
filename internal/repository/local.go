package repository

import "os"

type LocalRepository struct {
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{}
}

func (l *LocalRepository) SaveBytesToFile(filepath string, fileBytes []byte) error {
	file, createErr := os.Create(filepath)
	if createErr != nil {
		return createErr
	}

	_, writeErr := file.Write(fileBytes)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

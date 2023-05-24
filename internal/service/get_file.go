package service

import (
	"errors"
	"io"
	"os"
	fp "path/filepath"
)

var (
	IncorrectMetaErr = errors.New("incorrect meta")
)

type GetFileService struct {
}

func NewGetFileService() *GetFileService {
	return &GetFileService{}
}

func (f *GetFileService) ReceiveAndSaveFileByChunks(fileReceiver ChunkReceiver, filepath string) error {
	md, mdErr := checkAndGetMeta(fileReceiver)
	if mdErr != nil {
		return mdErr
	}

	file, createErr := os.Create(fp.Join(filepath, md["filename"]))
	if createErr != nil {
		return createErr
	}
	defer file.Close()

	for {
		fileChunk, recvErr := fileReceiver.Receive()
		if errors.Is(recvErr, io.EOF) {
			break
		}
		if recvErr != nil {
			return recvErr
		}

		_, writeErr := file.Write(fileChunk)
		if writeErr != nil {
			return writeErr
		}
	}

	return nil
}

func checkAndGetMeta(fileReceiver ChunkReceiver) (map[string]string, error) {
	md, mdErr := fileReceiver.Meta()
	if mdErr != nil {
		return nil, IncorrectMetaErr
	}

	_, ok := md["filename"]
	if !ok {
		return nil, errors.New("no filename in meta")
	}

	return md, nil
}

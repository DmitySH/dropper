package service

import (
	"bytes"
	"dmitysh/dropper/internal/entity"
	"io"
	"log"
	"strconv"
)

type GetFileService struct {
}

func NewGetFileService() *GetFileService {
	return &GetFileService{}
}

func (f *GetFileService) ParseDropCode(dropCode string) (entity.DropCode, error) {
	code, convErr := strconv.Atoi(dropCode)
	if convErr != nil {
		return entity.DropCode{}, convErr
	}

	return entity.DropCode{
		SecretCode: code % 100,
		HostID:     code / 100,
	}, nil
}

func (f *GetFileService) ReceiveFileByChunks(fileReceiver ChunkReceiver) ([]byte, error) {
	fileData := bytes.Buffer{}

	for {
		fileChunk, recvErr := fileReceiver.Receive()
		if recvErr == io.EOF {
			break
		}
		if recvErr != nil {
			return nil, recvErr
		}

		size := len(fileChunk)
		log.Printf("received a chunk with size: %d\n", size)

		_, writeErr := fileData.Write(fileChunk)
		if writeErr != nil {
			return nil, writeErr
		}
	}

	return fileData.Bytes(), nil
}

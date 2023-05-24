package service

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
)

var (
	FileAlreadyRequestedErr = errors.New("file already requested")
	TransferFileErr         = errors.New("error during file transfer")
)

type SendFileService struct {
	secretCode    int
	fileDropMu    sync.Mutex
	fileDropped   bool
	fileChunkSize int
}

func NewSendFileService(fileChunkSize int) *SendFileService {
	secretCode := rand.Intn(maxSecretCode-minSecretCode+1) + minSecretCode
	return &SendFileService{
		fileChunkSize: fileChunkSize,
		secretCode:    secretCode,
	}
}

func (f *SendFileService) SendFileByChunks(filepath string, fileSender ChunkSender) error {
	f.fileDropMu.Lock()
	if f.fileDropped {
		f.fileDropMu.Unlock()
		return FileAlreadyRequestedErr
	} else {
		f.fileDropped = true
		f.fileDropMu.Unlock()
	}

	if sendErr := f.sendFile(filepath, fileSender); sendErr != nil {
		log.Println("can't send file:", sendErr)
		f.fileDropMu.Lock()
		f.fileDropped = false
		f.fileDropMu.Unlock()

		return TransferFileErr
	}

	return nil
}

func (f *SendFileService) sendFile(filepath string, fileSender ChunkSender) error {
	file, openErr := os.Open(filepath)
	if openErr != nil {
		return fmt.Errorf("can't open file: %w", openErr)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, f.fileChunkSize)

	for {
		n, readChunkErr := reader.Read(buf)
		if errors.Is(readChunkErr, io.EOF) {
			break
		}

		if readChunkErr != nil {
			return fmt.Errorf("can't read chunk: %w", readChunkErr)
		}

		if sendErr := fileSender.Send(buf[:n]); sendErr != nil {
			return fmt.Errorf("can't send chunk: %w", sendErr)
		}
	}

	return nil
}

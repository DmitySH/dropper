package service

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	minSecretCode = 10
	maxSecretCode = 99

	maxIncorrectCodes = 2
)

var (
	OpenFileErr             = errors.New("can't open requested file")
	FileAlreadyRequestedErr = errors.New("file already requested")
	TransferFileErr         = errors.New("error during file transfer")
)

type SendFileService struct {
	secretCode            int
	fileDropMu            sync.Mutex
	codeAttemptsMu        sync.Mutex
	fileDropped           bool
	fileChunkSize         int
	incorrectCodeAttempts int
}

func NewSendFileService(fileChunkSize int) *SendFileService {
	return &SendFileService{fileChunkSize: fileChunkSize}
}

func (f *SendFileService) GenerateAndGetDropCode() string {
	f.secretCode = rand.Intn(maxSecretCode-minSecretCode+1) + minSecretCode
	dropCode := strings.Split(getOutboundIP().String(), ".")[3] + strconv.Itoa(f.secretCode)

	return dropCode
}

func (f *SendFileService) CheckSecretCode(code int) (bool, bool) {
	f.fileDropMu.Lock()
	defer f.fileDropMu.Unlock()
	f.incorrectCodeAttempts++

	return f.secretCode == code, f.incorrectCodeAttempts > maxIncorrectCodes
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

	file, openErr := os.Open(filepath)
	if openErr != nil {
		f.resetFileDropped()
		log.Println("can't open requested file:", openErr)
		return OpenFileErr
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, f.fileChunkSize)

	for {
		n, readChunkErr := reader.Read(buf)
		if readChunkErr == io.EOF {
			break
		}

		if readChunkErr != nil {
			f.resetFileDropped()
			log.Println("can't read chunk to buffer:", readChunkErr)
			return TransferFileErr
		}

		if sendErr := fileSender.Send(buf[:n]); sendErr != nil {
			f.resetFileDropped()
			log.Println("can't send chunk:", sendErr)
			return TransferFileErr
		}
	}

	return nil
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func (f *SendFileService) resetFileDropped() {
	f.fileDropMu.Lock()
	f.fileDropped = false
	f.fileDropMu.Unlock()
}

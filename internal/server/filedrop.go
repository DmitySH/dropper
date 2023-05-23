package server

import (
	"context"
	"dmitysh/dropper/internal/filedrop"
	"dmitysh/dropper/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	IncorrectCodeErr = status.Error(codes.InvalidArgument, "incorrect code")
)

type FileDropServer struct {
	filedrop.UnimplementedFileDropServer
	fileTransferService   service.SendFile
	filepath              string
	StopCh                chan os.Signal
	incorrectCodeAttempts int
	mu                    sync.Mutex
}

func NewFileDropServer(fileTransferService service.SendFile, filepath string, stopCh chan os.Signal) *FileDropServer {
	return &FileDropServer{fileTransferService: fileTransferService, StopCh: stopCh, filepath: filepath}
}

func (f *FileDropServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (f *FileDropServer) GetFile(_ *emptypb.Empty, fileStream filedrop.FileDrop_GetFileServer) error {
	if checkSecretCodeErr := f.checkSecretCode(fileStream.Context()); checkSecretCodeErr != nil {
		return checkSecretCodeErr
	}
	log.Println("file requested")

	streamSender := filedrop.NewStreamSender(fileStream)
	if sendFileErr := f.fileTransferService.SendFileByChunks(f.filepath, streamSender); sendFileErr != nil {
		return status.Error(codes.Internal, sendFileErr.Error())
	}

	log.Println("file transferred")
	f.StopCh <- os.Interrupt

	return nil
}

func (f *FileDropServer) checkSecretCode(mdCtx context.Context) error {
	md, ok := metadata.FromIncomingContext(mdCtx)
	if !ok {
		return status.Error(codes.InvalidArgument, "no meta provided")
	}

	secretCodeMeta := md.Get("secret-code")
	if len(secretCodeMeta) != 1 {
		return IncorrectCodeErr
	}

	secretCode, parseErr := strconv.Atoi(secretCodeMeta[0])
	if parseErr != nil {
		return IncorrectCodeErr
	}

	correct, tooMuchAttempts := f.fileTransferService.CheckSecretCode(secretCode)

	if tooMuchAttempts {
		f.StopCh <- os.Interrupt
		return IncorrectCodeErr
	}
	if !correct {
		log.Println("request with incorrect code")
		return IncorrectCodeErr
	}

	return nil
}

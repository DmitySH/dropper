package command

import (
	"context"
	"dmitysh/dropper/internal/filedrop"
	"dmitysh/dropper/internal/service"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
)

var (
	IncorrectCodeErr = errors.New("code is incorrect")
	ConnectionErr    = errors.New("can't connect")
	ReceiveErr       = errors.New("error during file receiving")
)

const (
	pingTimeout = time.Second * 1
)

type getOptions struct {
	dropCode int
}

func NewGetCommand() *cobra.Command {
	var options getOptions

	cmd := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get shared file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(cmd, &options, args)
		},
	}

	//flags := cmd.Flags()
	//flags.Int(&options.appName, "app-name", "n", "app", "Type of service controller ")

	return cmd
}

func runGet(_ *cobra.Command, _ *getOptions, args []string) error {
	fileGetterService := service.NewGetFileService()

	dropCode, parseCodeErr := fileGetterService.ParseDropCode(args[0])
	if parseCodeErr != nil {
		log.Println(parseCodeErr)
		return IncorrectCodeErr
	}

	conn, createConnErr := createConn(dropCode.HostID)
	if createConnErr != nil {
		log.Println(createConnErr)
		return ConnectionErr
	}
	defer conn.Close()

	fileDropClient := filedrop.NewFileDropClient(conn)
	if pingErr := pingServer(fileDropClient); pingErr != nil {
		return IncorrectCodeErr
	}

	fileStream, getFileStreamErr := getFileStream(dropCode.SecretCode, fileDropClient)
	if getFileStreamErr != nil {
		log.Println(getFileStreamErr)
		return ReceiveErr
	}
	streamReceiver := filedrop.NewStreamReceiver(fileStream)
	fileBytes, receiveFileErr := fileGetterService.ReceiveFileByChunks(streamReceiver)
	if receiveFileErr != nil {
		return errors.New(status.Convert(receiveFileErr).Message())
	}

	log.Println("file received. size:", len(fileBytes))

	return nil
}

func pingServer(fileDropClient filedrop.FileDropClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	_, pingErr := fileDropClient.Ping(ctx, &empty.Empty{})

	return pingErr
}

func createConn(hostID int) (*grpc.ClientConn, error) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	port := viper.GetInt("SERVER_PORT")
	localNetID := viper.GetString("LOCAL_NET_ID")

	fullAddr := fmt.Sprintf("%s.%d:%d", localNetID, hostID, port)

	conn, dialErr := grpc.Dial(fullAddr, opts...)
	if dialErr != nil {
		return nil, fmt.Errorf("can't create connection: %w", dialErr)
	}

	return conn, nil
}

func getFileStream(secretCode int, fileDropClient filedrop.FileDropClient) (filedrop.FileDrop_GetFileClient, error) {
	md := metadata.New(map[string]string{"secret-code": strconv.Itoa(secretCode)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	fileStream, getFileErr := fileDropClient.GetFile(ctx, &empty.Empty{})
	if getFileErr != nil {
		return nil, getFileErr
	}

	return fileStream, nil
}

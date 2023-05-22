package command

import (
	"dmitysh/your-drop/internal/filedrop"
	"dmitysh/your-drop/internal/server"
	"dmitysh/your-drop/internal/server/grpcutils"
	"dmitysh/your-drop/internal/service"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"os"
)

type dropOptions struct {
}

func NewDropCommand() *cobra.Command {
	var options dropOptions

	cmd := &cobra.Command{
		Use:   "drop [OPTIONS]",
		Short: "Share file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDrop(cmd, &options, args)
		},
	}

	//flags := cmd.Flags()
	//flags.StringVarP(&options.appName, "app-name", "n", "app", "Type of service controller ")

	return cmd
}

func runDrop(_ *cobra.Command, _ *dropOptions, args []string) error {
	fileSenderService := service.NewSendFileService(viper.GetInt("CHUNK_SIZE"))
	dropCode := fileSenderService.GenerateAndGetDropCode()

	fileDropServer := server.NewFileDropServer(fileSenderService, args[0], make(chan os.Signal, 2))

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	filedrop.RegisterFileDropServer(grpcServer, fileDropServer)

	serverCfg := grpcutils.GRPCServerConfig{
		Host: viper.GetString("SERVER_HOST"),
		Port: viper.GetInt("SERVER_PORT"),
	}

	fmt.Println("your drop code:", dropCode)

	runSrvErr := grpcutils.RunAndShutdownServer(serverCfg, grpcServer, fileDropServer)
	if runSrvErr != nil {
		return fmt.Errorf("can't serve: %w", runSrvErr)
	}

	return nil
}

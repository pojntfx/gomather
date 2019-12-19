package cmd

import (
	math "github.com/pojntfx/gomather/pkg/proto/generated/proto"
	"github.com/pojntfx/gomather/pkg/svc"
	"github.com/spf13/cobra"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	Port string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(command *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", Port)
		if err != nil {
			log.Fatal("Server could not listen", rz.Err(err))
		}

		server := grpc.NewServer()

		reflection.Register(server)
		math.RegisterMathServer(server, &svc.Math{})
		log.Info("Server started on port", rz.String("port", Port))

		if err := server.Serve(listener); err != nil {
			log.Fatal("Server could not start", rz.Err(err))
		}
	},
}

func init() {
	StartCmd.Flags().StringVarP(&Port, "port", "p", ":30000", "Server's port")

	RootCmd.AddCommand(StartCmd)
}

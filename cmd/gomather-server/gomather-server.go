package main

import (
	rz "gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
	"net"

	math "github.com/pojntfx/gomather/src/proto/generated/proto"
	"github.com/pojntfx/gomather/src/svc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

var rootCommand = &cobra.Command{
	Use:   "gomather-server",
	Short: "Simple Go gRPC microservice that does math.",
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(command *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal("Server could not listen", rz.Err(err))
		}

		server := grpc.NewServer()

		reflection.Register(server)
		math.RegisterMathServer(server, &svc.Math{})
		log.Info("Server started on port", rz.String("port", port))

		err = server.Serve(listener)
		if err != nil {
			log.Fatal("Server could not start", rz.Err(err))
		}
	},
}

func main() {
	startCommand.Flags().StringVarP(&port, "port", "p", ":30000", "Server's port")

	rootCommand.AddCommand(startCommand)

	err := rootCommand.Execute()
	if err != nil {
		log.Fatal("Server command could not start", rz.Err(err))
	}
}

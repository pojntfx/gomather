package main

import (
	log "github.com/sirupsen/logrus"
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
			log.Fatalln("Server could not listen", err)
		}

		server := grpc.NewServer()

		reflection.Register(server)
		math.RegisterMathServer(server, &svc.Math{})
		log.Println("Server started on port", port)

		err = server.Serve(listener)
		if err != nil {
			log.Fatalln("Server could not start", err)
		}
	},
}

func main() {
	startCommand.Flags().StringVarP(&port, "port", "p", ":30000", "Server's port")

	rootCommand.AddCommand(startCommand)

	err := rootCommand.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

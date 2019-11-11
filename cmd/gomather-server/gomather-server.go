package main

import (
	"flag"
	"log"
	"net"

	math "github.com/pojntfx/gomather/src/proto/generated/proto"
	"github.com/pojntfx/gomather/src/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// The port to run the gRPC server on
var port = flag.String("port", ":30000", "Server's port (by default :30000)")

func main() {
	// Parse the flags
	flag.Parse()

	// Start the listener
	listener, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalln("Server could not listen", err)
	}

	// Start the server
	server := grpc.NewServer()

	// Register the services
	reflection.Register(server)
	math.RegisterMathServer(server, &svc.Math{})
	log.Println("Server started on port", *port)

	// Serve the server via the listener
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Server could not start", err)
	}
}

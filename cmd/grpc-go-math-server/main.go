package main

import (
	"flag"
	"log"
	"net"

	"github.com/pojntfx/grpc-go-math/lib/proto"
	"github.com/pojntfx/grpc-go-math/lib/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// The port to run the gRPC server on
var port = flag.String("port", ":30000", "The port to run the gRPC server on")

func main() {
	// Parse the flags
	flag.Parse()

	// Start the listener
	listener, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalln("server could not listen", err)
	}

	// Start the server
	server := grpc.NewServer()

	// Register the services
	reflection.Register(server)
	proto.RegisterMathServer(server, &svc.Math{})
	log.Println("server started on port", *port)

	// Serve the server via the listener
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("server could not start", err)
	}
}

//go:generate protoc --go_out=paths=source_relative,plugins=grpc:../ -I=../ ../proto/math.proto

package svc

import (
	"context"
	"errors"
	"log"

	"github.com/pojntfx/grpc-go-math/lib/proto"
)

// Math is a gRPC example service
type Math struct {
	proto.UnimplementedMathServer
}

// Add adds two numbers
func (t *Math) Add(ctx context.Context, args *proto.MathAddArgs) (*proto.MathAddReply, error) {
	// Validate input
	if args.GetFirst() == 0 {
		return nil, errors.New("could not add, `First` has not been provided")
	}
	if args.GetSecond() == 0 {
		return nil, errors.New("could not add, `Second` has not been provided")
	}

	// Log progress
	log.Println("adding", args.GetFirst(), "to", args.GetSecond())

	// Return added numbers
	return &proto.MathAddReply{Result: args.GetFirst() + args.GetSecond()}, nil
}

// Subtract subtracts two numbers
func (t *Math) Subtract(ctx context.Context, args *proto.MathSubtractArgs) (*proto.MathSubtractReply, error) {
	// Validate input
	if args.GetFirst() == 0 {
		return nil, errors.New("could not subtract, `First` has not been provided")
	}
	if args.GetSecond() == 0 {
		return nil, errors.New("could not subtracts, `Second` has not been provided")
	}

	// Log progress
	log.Println("subtracting", args.GetSecond(), "from", args.GetFirst())

	// Return subtracted numbers
	return &proto.MathSubtractReply{Result: args.GetSecond() - args.GetFirst()}, nil
}

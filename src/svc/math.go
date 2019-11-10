//go:generate sh -c "mkdir -p ../proto/.generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/.generated --cobra_out=paths=source_relative,plugins=client:../proto/.generated -I=../ ../proto/math.proto"

package svc

import (
	"context"
	"errors"
	"log"

	"github.com/pojntfx/gomather/src/lib"
	"github.com/pojntfx/gomather/src/proto/.generated/math"
)

// Math is a gRPC example service
type Math struct {
	math.UnimplementedMathServer
}

// Add adds two numbers
func (t *Math) Add(ctx context.Context, args *math.MathAddArgs) (*math.MathAddReply, error) {
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
	return &math.MathAddReply{Result: lib.Add(args.GetFirst(), args.GetSecond())}, nil
}

// Subtract subtracts two numbers
func (t *Math) Subtract(ctx context.Context, args *math.MathSubtractArgs) (*math.MathSubtractReply, error) {
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
	return &math.MathSubtractReply{Result: lib.Subtract(args.GetFirst(), args.GetSecond())}, nil
}

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../ ../proto/mather.proto"

package svc

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/pojntfx/gomather/src/lib"
	math "github.com/pojntfx/gomather/src/proto/generated/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Math is a gRPC example service
type Math struct {
	math.UnimplementedMathServer
}

// Add adds two numbers
func (t *Math) Add(ctx context.Context, args *math.MathAddArgs) (*math.MathAddReply, error) {
	// Validate input
	if args.GetFirst() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Could not add, `First` has not been provided")
	}
	if args.GetSecond() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Could not add, `Second` has not been provided")
	}

	// Log progress
	log.Println("Adding", args.GetFirst(), "to", args.GetSecond())

	// Return added numbers
	return &math.MathAddReply{Result: lib.Add(args.GetFirst(), args.GetSecond())}, nil
}

// Subtract subtracts two numbers
func (t *Math) Subtract(ctx context.Context, args *math.MathSubtractArgs) (*math.MathSubtractReply, error) {
	// Validate input
	if args.GetFirst() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Could not add, `First` has not been provided")
	}
	if args.GetSecond() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Could not add, `Second` has not been provided")
	}

	// Log progress
	log.Println("Subtracting", args.GetSecond(), "from", args.GetFirst())

	// Return subtracted numbers
	return &math.MathSubtractReply{Result: lib.Subtract(args.GetFirst(), args.GetSecond())}, nil
}

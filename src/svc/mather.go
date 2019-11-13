//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../ ../proto/mather.proto"

package svc

import (
	"context"
	rz "gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"

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
		msg := "Could not add, `First` has not been provided"
		log.Error(msg)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}
	if args.GetSecond() == 0 {
		msg := "Could not add, `Second` has not been provided"
		log.Error(msg)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}

	// Log progress
	log.Info("Adding", rz.Int64("First", args.GetFirst()), rz.Int64("Second", args.GetSecond()))

	// Return added numbers
	return &math.MathAddReply{Result: lib.Add(args.GetFirst(), args.GetSecond())}, nil
}

// Subtract subtracts two numbers
func (t *Math) Subtract(ctx context.Context, args *math.MathSubtractArgs) (*math.MathSubtractReply, error) {
	// Validate input
	if args.GetFirst() == 0 {
		msg := "Could not subtract, `First` has not been provided"
		log.Error(msg)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}
	if args.GetSecond() == 0 {
		msg := "Could not subtract, `Second` has not been provided"
		log.Error(msg)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}

	// Log progress
	log.Info("Subtracting", rz.Int64("First", args.GetFirst()), rz.Int64("Second", args.GetSecond()))

	// Return subtracted numbers
	return &math.MathSubtractReply{Result: lib.Subtract(args.GetFirst(), args.GetSecond())}, nil
}

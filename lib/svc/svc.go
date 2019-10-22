//go:generate bash -c "protoc --go_out=plugins=grpc,paths=source_relative:../ --gqlgencfg_out=paths=source_relative:../ --gql_out=svcdir=true,paths=source_relative:../ --gogqlgen_out=paths=source_relative,gogoimport=false:../ -I=../ -I=/usr/local/include -I=`go list -m -f \"{{.Dir}}\" github.com/danielvladco/go-proto-gql/pb` ../proto/math.proto"

package svc

import (
	"context"
	"errors"
	"log"

	math "github.com/pojntfx/grpc-go-math/lib/proto"
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
	return &math.MathAddReply{Result: args.GetFirst() + args.GetSecond()}, nil
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
	return &math.MathSubtractReply{Result: args.GetSecond() - args.GetFirst()}, nil
}

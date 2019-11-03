# Felix Pojtinger's `grpc-go` Math Microservice

A simple math GRPC microservice, for the purpose of learning and evaluating Go and `grpc-go`.

## Features

- Add numbers
- Subtract numbers
- [Server as binary](./cmd/grpc-go-math-server/main.go)
- [Client as binary](./cmd/grpc-go-math-client/main.go)
- [Service as library](./lib/svc/svc.go)

## Usage

Binaries are made available on the [releases page](https://github.com/pojntfx/grpc-go-math/releases/latest). Alternatively, run the commands below to install from source:

```bash
# Install the proto compilers
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/fiorix/protoc-gen-cobra
# Compile the from proto
go generate ./...
# Download dependencies
go get ./...
# Build
go build ./...
# Install
go install ./...
# Run
grpc-go-math-server
```

## License

`grpc-go` Math Microservice (c) 2019 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0

# Felicitas Pojtinger's `grpc-go` Math Microservice

A simple math GRPC microservice, for the purpose of learning and evaluating Go and `grpc-go`.

## Features

- Add numbers
- Subtract numbers
- [Server as binary](./cmd/grpc-go-math-server/main.go)
- [Service as library](./lib/svc/svc.go)

## Usage

```bash
# Download
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/pojntfx/grpc-go-math/...
# Generate proto
go generate github.com/pojntfx/grpc-go-math/...
# Install
go install github.com/pojntfx/grpc-go-math/cmd/...
# Run
grpc-go-math-server
```

## License

`grpc-go` Math Microservice (c) 2019 Felicitas Pojtinger

SPDX-License-Identifier: AGPL-3.0

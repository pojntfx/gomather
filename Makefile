all: build

build:
	protoc -I ./ --go_out=plugins=grpc:. ./*/*/*.proto
	go build ./...

install:
	go install ./...
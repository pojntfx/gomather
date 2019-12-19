# syntax=docker/dockerfile:experimental
FROM --platform=$TARGETPLATFORM golang:1.13.5-buster AS build
WORKDIR /app
ARG TARGETPLATFORM

RUN apt update
RUN apt install -y protobuf-compiler

RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN curl -Lo dibs https://github.com/pojntfx/dibs/releases/latest/download/dibs-linux-amd64
RUN chmod +x dibs
RUN mv dibs /usr/local/bin/dibs

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./.dibs.yml ./.dibs.yml
COPY ./main.go ./main.go
COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN dibs pipeline build assets

FROM --platform=$TARGETPLATFORM debian:buster-slim
ARG TARGETPLATFORM

COPY --from=build /app/.bin/gomather-* /usr/local/bin/gomather

EXPOSE 30000

CMD gomather start
# syntax=docker/dockerfile:experimental
# Build container
FROM --platform=$TARGETPLATFORM golang:alpine AS build
ARG DIBS_TARGET
ARG TARGETPLATFORM

WORKDIR /app

RUN apk add -u curl protoc

RUN curl -Lo /tmp/dibs https://github.com/pojntfx/dibs/releases/latest/download/dibs-linux-amd64
RUN install /tmp/dibs /usr/local/bin

ENV GO111MODULE=on

RUN go get github.com/golang/protobuf/protoc-gen-go

ADD . .

RUN dibs -generateSources
RUN dibs -build

# Run container
FROM --platform=$TARGETPLATFORM alpine
ARG DIBS_TARGET
ARG TARGETPLATFORM

COPY --from=build /app/.bin/binaries/gomather* /usr/local/bin/gomather

CMD /usr/local/bin/gomather start

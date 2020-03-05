# gomather

Simple Go gRPC microservice that does math.

[![pipeline status](https://gitlab.com/pojntfx/gomather/badges/master/pipeline.svg)](https://gitlab.com/pojntfx/gomather/commits/master)

## Overview

gomather is an example gRPC microservice written in Go. It does math operations.

## Installation

### Prebuilt Binaries

Prebuilt binaries are available on the [releases page](https://github.com/pojntfx/gomather/releases/latest).

### Go Package

A Go package [is available](https://pkg.go.dev/github.com/pojntfx/gomather).

### Docker Image

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/gomather).

### Helm Chart

A Helm chart is available in [@pojntfx's Helm chart repository](https://pojntfx.github.io/charts/).

## Usage

```bash
% gomather
Simple Go gRPC microservice that does math.

Usage:
  gomather [command]

Available Commands:
  help        Help about any command
  start       Starts the server

Flags:
  -h, --help      help for gomather
      --version   version for gomather

Use "gomather [command] --help" for more information about a command.
```

## License

gomather (c) 2020 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0

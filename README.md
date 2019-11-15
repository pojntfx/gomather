# Mather Service (in Go)

Simple Go gRPC microservice that does math.

[![Build Status](https://travis-ci.com/pojntfx/gomather.svg?branch=master)](https://travis-ci.com/pojntfx/gomather)

## Features

- Add numbers
- Subtract numbers

## Usage

### From Prebuilt Binaries

Prebuilt binaries are available on the [releases page](https://github.com/pojntfx/gomather/releases/latest).

### From Go

```bash
# Install
go get -u github.com/pojntfx/gomather/cmd/gomather-server

# Start
gomather-server start
```

### From Source

#### Prerequisites

```bash
# Install dependencies
go get -u github.com/magefile/mage

# Install dependencies (for `protoc`)
PLATFORM=linux ARCHITECTURE=amd64 mage protocDependencyInstall
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage protocDependencyInstall

# Clean (optional)
mage clean

# Build
mage build
go get ./...
```

#### Start With Toolchain

```bash
# Start
mage start
```

#### Start As Standalone Binary

```bash
# Build binary
PLATFORM=linux ARCHITECTURE=amd64 mage binaryBuild
# or
PLATFORM=linux ARCHITECTURE=arm64 mage binaryBuild
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage binaryBuild

# Install binary
sudo -E env "PATH=$PATH" PLATFORM=linux ARCHITECTURE=amd64 mage binaryInstall
# or
sudo -E env "PATH=$PATH" PLATFORM=linux ARCHITECTURE=arm64 mage binaryInstall
# or
sudo -E env "PATH=$PATH" PLATFORM=darwin ARCHITECTURE=amd64 mage binaryInstall

# Start
gomather-server start
```

#### Unit Tests

```bash
# Start unit tests
mage unitTests
```

#### Integration Tests

```bash
# Start integration tests
mage integrationTests
```

#### Integration Tests (For Standalone Binary)

```bash
# Start integration tests (for standalone binary)
sudo -E env "PATH=$PATH" PLATFORM=linux ARCHITECTURE=amd64 mage binaryIntegrationTests
```

#### Development

```bash
# Start unit tests, start server and restart both if source changed
PLATFORM=linux ARCHITECTURE=amd64 mage dev
# or
PLATFORM=linux ARCHITECTURE=arm64 mage dev
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage dev
```

## License

Mather Service (in Go) (c) 2019 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0

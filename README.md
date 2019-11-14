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

```
# Install
go get -u github.com/pojntfx/gomather/cmd/gomather-server

# Run
gomather-server start
```

### From Source (Toolchain)

```bash
# Install dependencies
go get -u github.com/magefile/mage
PLATFORM=linux ARCHITECTURE=amd64 mage protocInstallDependencies

# Clean (optional)
mage clean

# Build
mage build
go get ./...

# Run
mage run
```

### From Source (Binary)

```bash
# Install dependencies
go get -u github.com/magefile/mage

# Install dependencies (for `protoc`)
PLATFORM=linux ARCHITECTURE=amd64 mage protocInstallDependencies
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage protocInstallDependencies

# Clean (optional)
mage clean

# Build
mage build
go get ./...

# Build binary
PLATFORM=linux ARCHITECTURE=amd64 mage binaryBuild
# or
PLATFORM=linux ARCHITECTURE=arm64 mage binaryBuild
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage binaryBuild

# Install binary
sudo PLATFORM=linux ARCHITECTURE=amd64 `which mage` binaryInstall
# or
sudo PLATFORM=linux ARCHITECTURE=arm64 `which mage` binaryInstall
# or
sudo PLATFORM=darwin ARCHITECTURE=amd64 `which mage` binaryInstall

# Run
gomather-server start
```

### From Source (Development)

```
# Install dependencies
go get -u github.com/magefile/mage

# Install dependencies (for `protoc`)
PLATFORM=linux ARCHITECTURE=amd64 mage protocInstallDependencies
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage protocInstallDependencies

# Watch, run and reload
PLATFORM=linux ARCHITECTURE=amd64 mage watch
# or
PLATFORM=linux ARCHITECTURE=arm64 mage watch
# or
PLATFORM=darwin ARCHITECTURE=amd64 mage watch
```

## License

Mather Service (in Go) (c) 2019 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0

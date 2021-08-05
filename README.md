# 📡 radar
A utility for monitoring SSL certs

## Usage

```shell
# Scan and output relevant SSL information for a host
 ./bin/radar scan <host>

```

## Contributing

Welcome to 📡 Radar, and thank you for contributing.

### Building the project

Prerequisites:

- Go 1.16

Run the following command to build:

```shell
# fetch and resolve dependencies
go mod tidy 

# build the binary
go build -o bin/radar ./cmd/radar/main.go

```

Run the binary:

- Unix-like systems: bin/gh
- Windows: bin\gh

Run tests with (What tests...?):

```shell
go test ./...
```

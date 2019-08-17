# gRPC Gateway

testing out [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway) with GoLang and TLS client authentication.

## Certificates

The certificates needed for this setup to work, is not provided in the repo, but see the [README](./certificates/README.md) for instructions on how to generate your own CA certificates, and server and client certificates in the `certificates` folder.

All certificates generated must be generated or placed in the `certificates` folder for this example to work.

## Use with Docker

The easiest way for using and testing this library is with Docker. This requires no other dependencie than docker on your local machine.

### Build image for generation
from the root dir run the following command:

    $ docker build -t go-grpc-build .

This will build an image with all tyhe depencies needed for building the requires files in this repository.

### Generate files
After building the `go-grpc-build` docker image, we can use this to generate the necessary files.

*Remember to set `TARGET` to your sitribution `linux`, `darwin` og `windows`*

    $ docker run --rm -v ${PWD}:/app -e TARGET=darwin go-grpc-build make

This will generate all protobuf files and binaries needed for running the test servers and clients.

**Remember to generate cerificates before testing**

### Run the binaries
#### gRPC Server

For any of this to work, the gPRC server must be running, as this is the one receiving all the request, either from the gRPC client, or through the HTTP proxy created by the grpc-gateway.

    $ cd server
    $ ./server.out

#### gRPC Client

The test gRPC client only requests one of the gRPC endpoints before shutting down. Run this in a separate terminal

    $ cd client
    $ ./client.out

#### HTTP Proxy client

The HTTP Proxy exposes an HTTP endpoint and talks with the gRPC gateway using the gateway definition file generated earlier. The client is set up so that a valid certificate must be presented also when using the REST api.

    $ cd httpServer
    $ ./httpserver.out

After this you can use [Postman](https://www.getpostman.com) to use the HTTP endpoints. For this to work, a valid certificate and key with the same CA cert authority as the gRPC server must be provided. For testing you may use the `client.crt` and `client.key` for this as well.
When using Postman the setting `SSL Certificate Verification` must be switched off for self-signed CA certs.

## Native Installation

The best way to test and use this repo is by following the docker description above. If you want to make this work on your local system, then keep on reading :)

The grpc-gateway requires a local installation of the Google protocol buffers compiler protoc v3.0.0 or above. This must be installed

Then get the golang requirements

    $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    $ go get -u github.com/golang/protobuf/protoc-gen-go

This will place three binaries in your `$GOBIN`;

- protoc-gen-grpc-gateway
- protoc-gen-swagger
- protoc-gen-go

Make sure that your `$GOBIN` is in your `$PATH`

### Usage

A simple `.proto` file is found in the `api/`folder, that contains one gRPC Service with two endpoints. There is also a `.yaml`file in `api/` that contains the information needed for gRPC-gateway. This is done like this to ned polute the `.proto` file with HTTP Proxy gateway stuff.

The generated files will not be included in this repo, so generating the source files must be done before starting.

Generate files for GO is done with make. To generate all use:

    $ make

To generate only the protobuff definition files:

    $ make proto

To generate only the binary files:

    $ make binaries

### Running

#### gRPC Server

For any of this to work, the gPRC server must be running, as this is the one receiving all the request, either from the gRPC client, or through the HTTP proxy created by the grpc-gateway.

    $ cd server
    $ go run server.go

Alternatively run the generated binary

    $ ./server.out

#### gRPC Client

The test gRPC client only requests one of the gRPC endpoints before shutting down. Run this in a separate terminal

    $ cd client
    $ go run client.go

Alternatively run the generated binary

    $ ./client.out

#### HTTP Proxy client

The HTTP Proxy exposes an HTTP endpoint and talks with the gRPC gateway using the gateway definition file generated earlier. The client is set up so that a valid certificate must be presented also when using the REST api.

    $ cd httpServer
    $ go run httpserver.go

Alternatively run the generated binary

    $ ./httpserver.out

After this you can use [Postman](https://www.getpostman.com) to use the HTTP endpoints. For this to work, a valid certificate and key with the same CA cert authority as the gRPC server must be provided. For testing you may use the `client.crt` and `client.key` for this as well.
When using Postman the setting `SSL Certificate Verification` must be switched off for self-signed CA certs.

## Summary

- Now all gRPC communication and HTTP endpoints are sucre and using TLS authentication for autorization. 
- No HTTP specific code is needed because everything is generated from the `.proto` and `.yaml` files provided in the `api/` folder.
- This also generates a `*.swagger.json` file that can be used with Swagger UI for interactive api documentation

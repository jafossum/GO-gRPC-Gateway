# gRPC Gateway

testing out [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway) with GoLang and TLS client authentication.

## Installation

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

## Usage

A simple `.proto` file is found in the `api/`folder, that contains one gRPC Service with two endpoints. There is also a `.yaml`file in `api/` that contains the information needed for gRPC-gateway. This is done like this to ned polute the `.proto` file with HTTP Proxy gateway stuff.

The generated files will not be included in this repo, so generating the source files must be done before starting.

Generate files for GO is done with make. To generate all use:

    $ make

To generate only the protobuff definition files:

    $ make proto

To generate only the gateway definition files:

    $ make gateway

To generate only the swagger definition files:

    $ make swagger

## Certificates

The certificates needed for this setup to work, is not provided in the repo, but see the [README](./certificates/README.md) for instructions on how to generate your own CA certificates, and server and client certificates in the `certificates` folder.

All certificates generated must be generated or placed in the `certificates` folder for this example to work.

## Running

### gRPC Server

For any of this to work, the gPRC server must be running, as this is the one receiving all the request, either from the gRPC client, or through the HTTP proxy created by the grpc-gateway.

    $ cd server
    $ go run server.go

### gRPC Client

The test gRPC client only requests one of the gRPC endpoints before shutting down. Run this in a separate terminal

    $ cd client
    $ go run client.go

### HTTP Proxy client

The HTTP Proxy exposes an HTTP endpoint and talks with the gRPC gateway using the gateway definition file generated earlier. The client is set up so that a valid certificate must be presented also when using the REST api.

    $ cd httpServer
    $ go run httpserver.go

After this you can use [Postman](https://www.getpostman.com) to use the HTTP endpoints. For this to work, a valid certificate and key with the same CA cert authority as the gRPC server must be provided. For testing you may use the `client.crt` and `client.key` for this as well.
When using Postman the setting `SSL Certificate Verification` must be switched off for self-signed CA certs.

### Summary

- Now all gRPC communication and HTTP endpoints are sucre and using TLS authentication for autorization. 
- No HTTP specific code is needed because everythong is generated from the `.proto` and `.yaml` files provided in the `api/` folder.
- This also generates a `*.swagger.json` file that can be used with Swagger UI for interactive api documentation

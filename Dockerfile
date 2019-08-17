FROM golang:1.12

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
 && rm -rf /var/lib/apt/lists/*

RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    && go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /app

COPY . .

RUN go mod download

FROM golang:1.12

# Install dependencies
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
 && rm -rf /var/lib/apt/lists/*

# Install GO dependencies
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    && go get -u github.com/golang/protobuf/protoc-gen-go

# Copy fiels to workdir
WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

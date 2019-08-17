# Go parameters
SERVICE=./api/test-service

.PHONY: binaries

all: proto binaries

proto:
	protoc -I /usr/local/include -I . -I ${GOPATH}/src --go_out=plugins=grpc:. ${SERVICE}.proto
	protoc -I /usr/local/include -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=${SERVICE}.yaml:. ${SERVICE}.proto
	protoc -I /usr/local/include -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --swagger_out=logtostderr=true,grpc_api_configuration=${SERVICE}.yaml:. ${SERVICE}.proto

binaries:
	cd ./server && GOOS=${TARGET} go build -o server.out && cd ..
	cd ./client && GOOS=${TARGET} go build -o client.out && cd ..
	cd ./httpServer && GOOS=${TARGET} go build -o httpServer.out && cd ..

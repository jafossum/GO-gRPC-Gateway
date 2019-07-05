# Go parameters
SERVICE=./api/test-service

.PHONY: all proto gateway

all: proto gateway swagger

proto:
	protoc -I /usr/local/include -I . -I ${GOPATH}/src --go_out=plugins=grpc:. ${SERVICE}.proto

gateway:
	protoc -I /usr/local/include -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=${SERVICE}.yaml:. ${SERVICE}.proto

swagger:
	protoc -I /usr/local/include -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --swagger_out=logtostderr=true,grpc_api_configuration=${SERVICE}.yaml:. ${SERVICE}.proto
.PHONY: run-json-rpc
run-json-rpc:
	go run ./cmd/json-rpc ./configs/main.yaml

.PHONY: run-grpc
run-grpc:
	go run ./cmd/grpc

.PHONY: build
build:
	go build -o ./build/app-skeleton-grpc ./cmd/grpc
	go build -o ./build/app-skeleton-json-rpc ./cmd/json-rpc

.PHONY: proto
proto:
	protoc --go_out=./api/grpc/ --go-grpc_out=./api/grpc/ api/grpc/proto/*.proto

.PHONY: install-ubuntu
install-ubuntu:
	apt install -y protobuf-compiler
	protoc --version

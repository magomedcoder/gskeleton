.PHONY: install
install:
	go install github.com/google/wire/cmd/wire@latest \
	&& go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	&& go install github.com/srikrsna/protoc-gen-gotag@latest

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
	protoc --proto_path api/grpc \
		   --go_out=pkg/pb_generated \
		   --go_opt=paths=source_relative \
		   --go-grpc_opt=paths=source_relative \
		   --go-grpc_out=pkg/pb_generated \
		   api/grpc/*.proto

.PHONY: install-ubuntu
install-ubuntu:
	apt install -y protobuf-compiler
	protoc --version

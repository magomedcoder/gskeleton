.PHONY: install
install:
	go install github.com/google/wire/cmd/wire@latest \
	&& go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	&& go install github.com/srikrsna/protoc-gen-gotag@latest

.PHONY: run-http
run-http:
	go run ./cmd/gskeleton run-http -config ./configs/gskeleton.yaml

.PHONY: run-grpc
run-grpc:
	go run ./cmd/gskeleton run-grpc -config ./configs/gskeleton.yaml

.PHONY: build
build:
	go build -o ./build/gskeleton ./cmd/gskeleton

.PHONY: proto
proto:
	protoc --proto_path ./api/grpc/proto \
		   --go_out=paths=source_relative:./api/grpc/pb \
		   --go-grpc_out=paths=source_relative:./api/grpc/pb \
		   ./api/grpc/proto/*.proto


.PHONY: install-ubuntu
install-ubuntu:
	apt install -y protobuf-compiler
	protoc --version

.PHONY: install-ubuntu
install-ubuntu:
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v30.2/protoc-30.2-linux-x86_64.zip
	unzip protoc-30.2-linux-x86_64.zip -d $HOME/.local
	export PATH="$PATH:$HOME/.local/bin"
	protoc --version

.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	&& go install github.com/srikrsna/protoc-gen-gotag@latest

.PHONY: run-http
run-http:
	go run ./cmd/gskeleton run http -config ./configs/gskeleton.yaml

.PHONY: run-grpc
run-grpc:
	go run ./cmd/gskeleton run grpc -config ./configs/gskeleton.yaml

.PHONY: cli-migrate
cli-migrate:
	go run ./cmd/gskeleton cli migrate -config ./configs/gskeleton.yaml

.PHONY: cli-create-user
cli-create-user:
	go run ./cmd/gskeleton cli create-user -config ./configs/gskeleton.yaml

test:
	go test -v ./...

.PHONY: build
build:
	go build -o ./build/gskeleton ./cmd/gskeleton

.PHONY: gen
gen:
	protoc --proto_path=./api/grpc/proto \
	   --go_out=paths=source_relative:./api/grpc/pb \
	   --go-grpc_out=paths=source_relative:./api/grpc/pb \
	   ./api/grpc/proto/*.proto

	protoc --proto_path=./api/http/proto \
	  --proto_path=./third_party/proto/ \
	  --go_out=paths=source_relative:./api/http/pb \
	  --validate_out=paths=source_relative,lang=go:./api/http/pb/ \
	  ./api/http/proto/v1/*.proto

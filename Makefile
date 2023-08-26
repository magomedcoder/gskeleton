.PHONY: run
run:
	go run ./cmd/rpc ./configs/main.yaml

.PHONY: build
build:
	go build -o ./build/rpc-server ./cmd/rpc

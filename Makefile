.PHONY: run
run:
	go run ./cmd/app ./configs/main.yaml

.PHONY: build
build:
	go build -o ./build/app-skeleton ./cmd/app

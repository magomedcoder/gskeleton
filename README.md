## Golang app skeleton

## Install

```shell
go mod tidy
```

## Run

```shell
make run
```

## Build

```shell
make build
```

## Folder structure

```
├── cmd
│    └── app
│        ├── main.go
│        ├── wire.go
│        └── wire_gen.go
├── configs
│   └── main.yaml
├── internal
│   ├── config
│   │   └── config.go
│   ├── provider
│   │   └── server.go
│   └── transport
│       └── app
│           ├── handler
│           │   ├── example.go
│           │   └── handler.go
│           └── router
│               └── router.go
├── pkg
│   └── server
│       ├── error.go
│       ├── http.go
│       ├── options.go
│       ├── rpc.go
│       ├── server.go
│       └── transport.go
├── .editorconfig
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

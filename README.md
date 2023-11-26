## Golang app skeleton

## Install

```shell
go mod tidy
```

## Run

```shell
make run
```

## Call example

```
{
    "jsonrpc": "2.0",
    "method": "example.get",
    "params": {},
    "id": 1
}

{
    "jsonrpc": "2.0",
    "method": "example.set",
    "params": {
        "text": "Test"
    },
    "id": 1
}
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

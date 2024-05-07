## Golang app skeleton

## Folder structure

```
├── cmd
│    ├── grpc
│    │   ├── main.go
│    │   ├── wire.go
│    │   └── wire_gen.go
│    └── json-rpc
│        ├── main.go
│        ├── wire.go
│        └── wire_gen.go
├── configs
│   └── main.yaml
├── internal
│   ├── config
│   │   ├── config.go
│   │   ├── jwt.go
│   │   ├── postgres.go
│   │   └── server.go
│   ├── provider
│   │   ├── grpc_server.go
│   │   ├── json_rpc_server.go
│   │   └── postgres.go
│   └── transport
│   │   ├── model
│   │   │   └─ user.go
│   │   └── repo
│   │       └─ user.go
│   └── transport
│       ├── grpc
│       │   ├── handler
│       │   │   ├── auth.go
│       │   │   └── user.go
│       │   ├── middleware
│       │   │   ├── auth.go
│       │   │   ├── global.go
│       │   │   └── token.go
│       │   └── router
│       │       └── methods.go
│       └── json-rpc
│           ├── handler
│           │   ├── example.go
│           │   └── handler.go
│           └── router
│               └── router.go
├── pkg
│   └── json-rpc-server
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

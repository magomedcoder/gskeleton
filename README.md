## Golang app skeleton

## Folder structure

```
├─ api
│   └─ grpc
│       ├─ auth.proto
│       └─ user.proto
├─ cmd
│   ├─ grpc
│   │   ├─ main.go
│   │   ├─ wire.go
│   │   └─ wire_gen.go
│   │
│   ├─ http
│   │   ├─ main.go
│   │   ├─ wire.go
│   │   └─ wire_gen.go
│   │
│   └── json-rpc
│        ├─ main.go
│        ├─ wire.go
│        └─ wire_gen.go
│
├─ configs
│   └─ main.yaml
│
├─ internal
│   ├─ config
│   │   ├─ app.go
│   │   ├─ jwt.go
│   │   ├─ postgres.go
│   │   └─ server.go
│   │
│   ├─ model
│   │   └─ error.go
│   │
│   ├─ provider
│   │   ├─ app.go
│   │   ├─ grpc.go
│   │   ├─ json_rpc.go
│   │   └─ postgres.go
│   │
│   ├─ repository
│   │   ├─ user
│   │   │   ├─ entity
│   │   │   │   └─ user.go
│   │   │   │
│   │   │   └─ repo
│   │   │       ├─ user.go
│   │   │       ├─ user_create.go
│   │   │       └─ user_get.go
│   │   │
│   │   └─ repository.go
│   │
│   ├─ service
│   │   ├─ user
│   │   │   └─ user.go
│   │   └─ postgres.go
│   │
│   └─ transport
│       ├─ grpc
│       │   ├─ handler
│       │   │   ├─ auth.go
│       │   │   └─ user.go
│       │   │
│       │   ├─ middleware
│       │   │   ├─ auth.go
│       │   │   ├─ global.go
│       │   │   └─ token.go
│       │   │
│       │   └─ router
│       │       └─ methods.go
│       ├─ http
│       │   ├─ handler
│       │   │   └─ v1
│       │   │       ├─ example.go
│       │   │       └─ handler.go
│       │   │
│       │   ├─ middleware
│       │   │   ├─ auth.go
│       │   │   └─ handler.go
│       │   │
│       │   ├─ router
│       │   │   └─ methods.go
│       │   │
│       │   ├─ server.go
│       │   └─ wire.go
│       │
│       └─ json-rpc
│           ├─ handler
│           │   ├─ example.go
│           │   └─ handler.go
│           │
│           └─ router
│               └─ router.go
│
├─ pkg
│   ├─ jencrypt
│   │   └─ encrypt.go
│   │
│   ├─ http-server
│   │   ├─ context.go
│   │   ├─ handler.go
│   │   └─ response.go
│   │
│   └─ json-rpc-server
│   │   ├─ error.go
│   │   ├─ http.go
│   │   ├─ options.go
│   │   ├─ rpc.go
│   │   ├─ server.go
│   │   └─ transport.go
│   │
│   ├─ jsonutil
│   │   └─ json.go
│   │
│   └─ strutil
│       └─ str.go
├─ test
│   └─ rpc-call.http
│
├─ .editorconfig
├─ .gitignore
├─ docker-compose.yaml
├─ go.mod
├─ go.sum
├─ LICENSE
├─ Makefile
└─ README.md
```

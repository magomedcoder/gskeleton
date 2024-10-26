## Golang - The Ultimate Folder Structure

Organizing your Go (Golang) project's folder structure can help improve code readability, maintainability, and
scalability.
While there is no one-size-fits-all structure, here's a common folder structure for a Go project:

```
├─ api
│   └─ grpc
│       ├─ pb
│       └─ proto
│           ├─ auth.proto
│           └─ user.proto
├─ cmd
│   └── gskeleton
│        ├─ main.go
│        ├─ wire.go
│        └─ wire_gen.go
├─ configs
│   └─ gskeleton.yaml
├─ internal
│   ├─ config
│   │   ├─ app.go
│   │   ├─ jwt.go
│   │   ├─ postgres.go
│   │   └─ server.go
│   ├─ model
│   │   └─ error.go
│   ├─ provider
│   │   ├─ app.go
│   │   └─ postgres.go
│   ├─ repository
│   │   ├─ user
│   │   │   ├─ entity
│   │   │   │   └─ user.go
│   │   │   └─ repo
│   │   │       ├─ user.go
│   │   │       ├─ user_create.go
│   │   │       └─ user_get.go
│   │   └─ repository.go
│   ├─ service
│   │   ├─ user
│   │   │   └─ user.go
│   │   └─ postgres.go
│   └─ transport
│       ├─ grpc
│       │   ├─ handler
│       │   │   ├─ auth.go
│       │   │   └─ user.go
│       │   ├─ middleware
│       │   │   ├─ authorization.go
│       │   │   ├─ global.go
│       │   │   ├─ logging.go
│       │   │   └─ token.go
│       │   ├─ server.go
│       │   └─ wire.go
│       └─ http
│           ├─ handler
│           │   └─ v1
│           │       ├─ example.go
│           │       └─ handler.go
│           ├─ middleware
│           │   ├─ auth.go
│           │   └─ handler.go
│           ├─ router
│           │   └─ methods.go
│           ├─ server.go
│           └─ wire.go
├─ migrations
│   └─ postgres
│       ├─ 000_migration.down.sql
│       └─ 000_migration.up.sql
├─ pkg
│   ├─ jencrypt
│   │   └─ encrypt.go
│   ├─ http-server
│   │   ├─ context.go
│   │   ├─ handler.go
│   │   └─ response.go
│   ├─ jsonutil
│   │   └─ json.go
│   └─ strutil
│       └─ str.go
├─ test
│   └─ test-http.http
├─ .editorconfig
├─ .gitignore
├─ docker-compose.yaml
├─ go.mod
├─ go.sum
├─ LICENSE
├─ Makefile
└─ README.md
```

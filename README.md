# 1. Golang REST API Server

- [1. Golang REST API Server](#1-golang-rest-api-server)
    - [1.0.1. Useful links](#101-useful-links)
    - [1.0.2. Used libs](#102-used-libs)
    - [1.0.3. `go mod`](#103-go-mod)
      - [1.0.3.1. `go mod` commands](#1031-go-mod-commands)

### 1.0.1. Useful links

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Codeship Golang Best Practices](https://github.com/codeship/go-best-practices)
- [Go database/sql tutorial](http://go-database-sql.org/index.html)

### 1.0.2. Used libs

- [**TOML**: TOML parser for Golang with reflection.](https://github.com/BurntSushi/toml)
- [**logrus**: Structured, pluggable logging for Go.](https://github.com/sirupsen/logrus)
- [**mux**: A powerful HTTP router and URL matcher for building Go web servers with Gorilla](https://github.com/gorilla/mux)
- [**testify**: A toolkit with common assertions and mocks that plays nicely with the standard library](https://github.com/stretchr/testify)
- [**pq**: Pure Go Postgres driver for database/sql](https://github.com/lib/pq)

### 1.0.3. `go mod`

- `go mod` package manager for Golang,
- `go mod init github.com/ivantokar/repository` initialize new module in current directory, it creates a _go.mod_ file

#### 1.0.3.1. `go mod` commands

| command  | description                                |
| -------- | ------------------------------------------ |
| download | download modules to local cache            |
| edit     | edit go.mod from tools or scripts          |
| graph    | print module requirement graph             |
| init     | initialize new module in current directory |
| tidy     | add missing and remove unused modules      |
| vendor   | make vendored copy of dependencies         |
| verify   | verify dependencies have expected content  |
| why      | explain why packages or modules are needed |

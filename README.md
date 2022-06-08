# try-golang

This is my TUTORIAL project for golang

![try-golang - Go Version](https://img.shields.io/badge/go-1.18-blue.svg)
[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge/master)](https://www.codefactor.io/repository/github/devlights/try-golang/overview/master)
![Go](https://github.com/devlights/try-golang/workflows/Go/badge.svg?branch=master)

## Go version

```shell script
$ go version
go version go1.18 linux/amd64

$ make build
go build -race

$ go version try-golang
try-golang: go1.18
```

## Run

```shell script
$ go run main.go
```

or 

```shell script
$ make run
```

or

```shell script
$ make run EXAMPLE=example_name
```

or

```shell script
$ make docker
```

or

If you want to use [go-task](https://github.com/go-task/task), type the following command.

```sh
$ go install github.com/go-task/task/v3/cmd/task@latest
```

Once the above command is complete, you can run it at

```sh
$ task run
```

### Run Docker on Gitpod

type following command in first terminal:

```shell script
$ sudo docker-up
```

Launch new terminal and type following command:

```shell
$ make docker
```

## Test

```shell script
$ go test -v ./...
```

or

```shell script
$ make test
```

## Install

```shell script
$ go install
```

or 

```shell script
$ make install
```

## 関連リポジトリ

- [gomy](https://github.com/devlights/gomy)
  - 共通ライブラリ
- [try-golang-extlib](https://github.com/devlights/try-golang-extlib)
  - 3rd party ライブラリのサンプルはこちらで管理しています。
- [try-golang-db](https://github.com/devlights/try-golang-db)
  - データベースのサンプルはこちらで管理しています。
- [go-crosscompile-example](https://github.com/devlights/go-crosscompile-example)
  - Goでクロスコンパイルを行うサンプルです。
- [go-grpc-uds-example](https://github.com/devlights/go-grpc-uds-example)
  - GoでgRPCで unix domain socket を扱うサンプルです。
- [go-protobuf-example](https://github.com/devlights/go-protobuf-example)
  - Goで protocol buffers を扱うサンプルです。
- [go-unix-domain-socket-example](https://github.com/devlights/go-unix-domain-socket-example)
  - Go で unix domain socket を使って通信するサンプルです。
- [go-mod-vendoring-example](https://github.com/devlights/go-mod-vendoring-example)
  - Go で 依存しているモジュールを vendoring して実行してみるサンプルです。

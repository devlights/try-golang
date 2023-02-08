
# try-golang

This is my TUTORIAL project for golang

![try-golang - Go Version](https://img.shields.io/badge/go-1.19-blue.svg)
[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge/master)](https://www.codefactor.io/repository/github/devlights/try-golang/overview/master)
![Go](https://github.com/devlights/try-golang/workflows/Go/badge.svg?branch=master)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/devlights/try-golang)

## Go version

```shell script
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.5 LTS
Release:        20.04
Codename:       focal

$ go version
go version go1.19.2 linux/amd64

$ task build
task: [build] go build .

$ go version try-golang
try-golang: go1.19.2
```

## Run

```shell script
$ go run main.go
```

If you want to use [go-task](https://github.com/go-task/task), type the following command.

```sh
$ go install github.com/go-task/task/v3/cmd/task@latest
```

Once the above command is complete, you can run it at

```sh
$ task run
```

## Test

```shell script
$ go test -v ./...
```

or

```shell script
$ task test
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
- [go-listener-with-backlog-example](https://github.com/devlights/go-listener-with-backlog-example)
  - Go で バックログ を指定できる net.Listener を生成して実行するサンプルです。
- [go-syscall-with-strace](https://github.com/devlights/go-syscall-with-strace)
  - Go アプリが内部で利用しているシステムコールを出力するサンプルです。
- [go-socket-reuseport-example](https://github.com/devlights/go-socket-reuseport-example)
  - Go で ソケット の SO_REUSEPORT を有効にして、同一ポートに複数LISTENするサーバを作るサンプルです。

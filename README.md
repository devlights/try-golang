
# try-golang

This is my TUTORIAL project for golang.

![try-golang - Go Version](https://img.shields.io/badge/go-1.24-blue.svg)

## Run

```sh
$ go run ./cmd/try-golang
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
- [goxcel](https://github.com/devlights/goxcel)
  - Excel操作ライブラリ (go-ole利用)
- [gord](https://github.com/devlights/gord)
  - Word操作ライブラリ (go-ole利用)
- [try-golang-extlib](https://github.com/devlights/try-golang-extlib)
  - 3rd party ライブラリのサンプルはこちらで管理しています。
- [try-golang-db](https://github.com/devlights/try-golang-db)
  - データベースのサンプルはこちらで管理しています。
- [try-golang-cgo](https://github.com/devlights/try-golang-cgo)
  - CGOのサンプルはこちらで管理しています。
- [try-golang-network](https://github.com/devlights/try-golang-network)
  - ネットワークのサンプルはこちらで管理しています。
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
- [go-gopacket-example](https://github.com/devlights/go-gopacket-example)
  - [gopacket](https://github.com/google/gopacket)を使ってパケットキャプチャを行うサンプルです。
- [go-tcp-keepalive-example](https://github.com/devlights/go-tcp-keepalive-example)
  - Go で キープアライブプローブ を送信するサンプルです。
- [バイトスライスに文字列を速く設定する方法（fmt.Sprintf, fmt.Appendf, 直接append使用)](https://gist.github.com/devlights/ffd22f78297a563c9bebcb9a9baa7f5f)
- [go124-goget-tools-dependencies-example](https://github.com/devlights/go124-goget-tools-dependencies-example)
  - Go 1.24で導入された go get -tool によるツール依存関係インストールのサンプルです。

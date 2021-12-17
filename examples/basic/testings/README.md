# 概要

このパッケージのみ、```$ make run``` でサンプルを動かすことができません。

実行する場合は

```sh
$ cd examples/basic/testings
$ go test -v ./...
```

とするか

```sh
$ go test -v github.com/devlights/try-golang/examples/basic/testings/...
```

としてください。

## 参考

- https://medium.com/better-programming/easy-guide-to-unit-testing-in-golang-4fc1e9d96679
- https://github.com/golang/go/wiki/LearnTesting

# 実行方法

go-task経由で実行すると go-task がシグナルをハンドリングしてしまうので

本サンプルに関しては、直接実行する必要があります。

```sh
$ go run .
5秒間 SIGINT をハンドリング
^CCTRL-C PRESSED
^CCTRL-C PRESSED
^CCTRL-C PRESSED
signal.Reset() CALLED
^Csignal: interrupt
```

# これは何？

[golang.org/sys/unix](https://pkg.go.dev/golang.org/x/sys/unix) を使って、[select(2)](https://ja.manpages.org/select/2)を呼び出し

対象となるソケットファイルディスクリプタが読み取り可能な状態であるかを確認して処理するサンプルです。


```sh
$ task
task: [build] go build -o app .
task: [run] ./app -server &
task: [run] sleep 0.5
task: [run] ./app
10:01:00.206680 [C] select(2) -- not readable(fd=6)
10:01:00.216917 [C] select(2) -- not readable(fd=6)
10:01:00.227016 [C] select(2) -- not readable(fd=6)
10:01:00.237166 [C] select(2) -- not readable(fd=6)
10:01:00.247262 [C] select(2) -- not readable(fd=6)
10:01:00.257365 [C] select(2) -- not readable(fd=6)
10:01:00.267527 [C] select(2) -- not readable(fd=6)
10:01:00.277664 [C] select(2) -- not readable(fd=6)
10:01:00.287777 [C] select(2) -- not readable(fd=6)
10:01:00.296891 [S] send data
10:01:00.297322 [C] recv hello
10:01:00.297369 [C] shutdown(SHUT_WR)
10:01:00.297405 [C] close
10:01:00.297458 [S] disconnect
10:01:00.297523 [S] close
task: [run] sleep 0.5
task: [run] pkill app
```

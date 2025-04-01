# UNIXドメインソケットの抽象名前空間のサンプル (net-conn版)

Linux特有の機能である「抽象名前空間」（Abstract Namespace）のUnixドメインソケットは以下の特徴がある

- ソケットのアドレス（名前）の先頭にNULバイト（\0）を付ける
- ファイルシステム上にソケットファイルを作成しない
- プロセス終了時に自動的にクリーンアップされる

netパッケージで利用する場合、アドレスの先頭に @ を付与しておくと、内部で NULバイト(\0) に置き換えてくれる。

## 実行例

```sh
$ task
task: [build] go build -o app main.go
task: [run] ./app -server &
task: [run] sleep 1
task: [run] ./app
02:42:06.032066 [C] Send (hello)
02:42:06.032128 [S] Recv (hello)
02:42:06.032345 [S] Send (HELLO)
02:42:06.032365 [C] Recv (HELLO)
02:42:06.032378 [C] SEND FIN (shutdown(SHUT_WR))
02:42:06.032382 [S] disconnect
02:42:06.032410 [S] close
02:42:06.032413 [C] disconnect
02:42:06.032436 [C] close
task: [run] pkill -INT -f './app -server'
02:42:06.043557 [S] Shutdown...
```

## 参考情報

- https://man7.org/linux/man-pages/man7/unix.7.html
- https://siguniang.wordpress.com/2012/04/29/unix-domain-socket-address-types/

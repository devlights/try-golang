# UNIXドメインソケットの抽象名前空間のサンプル (net版)

Linux特有の機能である「抽象名前空間」（Abstract Namespace）のUnixドメインソケットは以下の特徴がある

- ソケットのアドレス（名前）の先頭にNULバイト（\0）を付ける
- ファイルシステム上にソケットファイルを作成する
- プロセス終了時に自動的にクリーンアップされる

netパッケージで利用する場合、アドレスの先頭に @ を付与しておくと、内部で NULバイト(\0) に置き換えてくれる。

## 実行例

```sh
$ task
task: [build] go build -o app main.go
task: [run] ./app -server &
task: [run] sleep 1
task: [run] ./app
09:01:13.238544 [C] Send (hello)
09:01:13.239925 [S] Recv (hello)
09:01:13.240143 [S] Send (HELLO)
09:01:13.240205 [C] Recv (HELLO)
09:01:13.240314 [S] disconnect
task: [run] pkill -f './app -server' || true
```

## 参考情報

- https://man7.org/linux/man-pages/man7/unix.7.html
- https://siguniang.wordpress.com/2012/04/29/unix-domain-socket-address-types/

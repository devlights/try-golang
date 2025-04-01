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
01:59:18.896796 [C] Send (hello)
01:59:18.896882 [S] Recv (hello)
01:59:18.897098 [S] Send (HELLO)
01:59:18.897149 [C] Recv (HELLO)
01:59:18.897174 [C] close
01:59:18.897177 [S] disconnect
01:59:18.897205 [S] close
task: [run] pkill -INT -f './app -server'
01:59:18.909751 [S] Shutdown...
```

## 参考情報

- https://man7.org/linux/man-pages/man7/unix.7.html
- https://siguniang.wordpress.com/2012/04/29/unix-domain-socket-address-types/

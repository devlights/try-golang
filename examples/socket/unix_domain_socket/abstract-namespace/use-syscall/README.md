# UNIXドメインソケットの抽象名前空間のサンプル (syscall版)

Linux特有の機能である「抽象名前空間」（Abstract Namespace）のUnixドメインソケットは以下の特徴がある

- ソケットのアドレス（名前）の先頭にNULバイト（\0）を付ける
- ファイルシステム上にソケットファイルを作成しない
- プロセス終了時に自動的にクリーンアップされる

syscallパッケージを利用した場合はC言語と同じルールで処理を実装することになる。

netパッケージでも同じことを実装することができる。

## 実行例

```sh
$ task
task: [build] go build -o app main.go
task: [run] ./app -server &
task: [run] sleep 1
task: [run] ./app
08:23:30.473291 [S] Connect from @
08:23:30.473439 [S] Recv (hello)
08:23:30.473446 [S] Send (HELLO)
08:23:30.473435 [C] Send (hello)
08:23:30.473604 [C] Recv (HELLO)
08:23:30.473618 [S] disconnect
task: [run] pkill -f './app -server' || true
```

## 参考情報

- https://man7.org/linux/man-pages/man7/unix.7.html
- https://siguniang.wordpress.com/2012/04/29/unix-domain-socket-address-types/

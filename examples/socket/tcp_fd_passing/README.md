# ファイルディスクリプタパッシングのサンプル

[sys/unix](https://pkg.go.dev/golang.org/x/sys/unix) を用いて、ファイルディスクリプタパッシングを行うサンプルです。

ファイルディスクリプタパッシング（FD passing）とは、あるプロセスから別のプロセスへ、既に開かれたファイルディスクリプタを転送する技術である。
Unixシステムでは、ファイル、ソケット、パイプなどのリソースは、プロセス内でファイルディスクリプタという整数値で表現される。

通常、各プロセスは独自のファイルディスクリプタテーブルを持つが、UnixドメインソケットのSCM_RIGHTS機能を使用することで
あるプロセスのファイルディスクリプタを別のプロセスに転送し、そのプロセスからも同じリソースにアクセスできるようにすることが出来る。

```sh
$ task
task: [build] go build -o tcp-client tcpclient/main.go
task: [build] go build -o tcp-server tcpserver/main.go
task: [build] go build -o uds-server udsserver/main.go
task: [run] ./uds-server &
task: [run] sleep 1
05:12:51.462037 [UDS-S] uds-server listening on
task: [run] ./tcp-server &
task: [run] sleep 1
05:12:52.475893 [UDS-S] @
05:12:52.475846 [TCP-S] connect uds-server
05:12:52.476473 [TCP-S] tcp-listen on :8888
task: [run] ./tcp-client
05:12:53.489672 [TCP-S] accept client
05:12:53.489720 [TCP-S] passing fd=8 to uds-server
05:12:53.489742 [TCP-S] close
05:12:53.489650 [TCP-C] connect tcp-server
05:12:53.489781 [UDS-S] recv fd=7
05:12:53.489868 [UDS-S] send (hello)
05:12:53.489894 [TCP-C] recv (hello)
05:12:53.489945 [TCP-C] send (HELLO)
05:12:53.489955 [UDS-S] recv (HELLO)
05:12:53.489978 [UDS-S] shutdown(SHUT_WR)
05:12:53.489984 [TCP-C] disconnect
05:12:53.489990 [UDS-S] close
05:12:53.490041 [TCP-C] close
task: [run] pkill tcp-server
task: [run] pkill uds-server
```

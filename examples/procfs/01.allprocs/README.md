# これは何？

[procfs](https://github.com/prometheus/procfs) を使って、/procファイルシステム上の全プロセス情報を取得するサンプルです。

```sh
$ task
task: [default] rm -f ./app
task: [default] go build -o app .
task: [default] ./app
[01]       1: /.supervisor/supervisor
[02]      36: /.supervisor/supervisor
[03]      70: /usr/bin/dash
[04]      78: /ide/node
[05]      90: /usr/bin/bash
[06]     666: /ide/node
[07]    1286: /ide/node
[08]    1952: /ide/node
[09]    1977: /ide/node
[10]    2492: /ide/node
[11]    7471: /home/gitpod/go-packages/bin/gopls
[12]    7483: /home/gitpod/go-packages/bin/gopls
[13]   23327: /ide/node
[14]   42093: /home/gitpod/go-packages/bin/staticcheck
[15]   42113: /home/gitpod/go/bin/go
[16]   42122: /workspace/go/bin/task
[17]   42439: /workspace/try-golang/examples/procfs/allprocs/app
```

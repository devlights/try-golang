# これは何？

[procfs](https://github.com/prometheus/procfs) を使って、/procファイルシステム上の自プロセス情報を取得するサンプルです。

```sh
$ task
task: [default] rm -f ./app
task: [default] go build -o app .
task: [default] ./app
[Self] pid=21966, cmdline=[./app]
task: [default] ./app hello world 1 2 3 4 5
[Self] pid=21972, cmdline=[./app hello world 1 2 3 4 5]
```

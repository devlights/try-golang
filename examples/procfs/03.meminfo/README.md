# これは何？

[procfs](https://github.com/prometheus/procfs) を使って、/procファイルシステム上のメモリ情報を取得するサンプルです。

```sh
$ task
task: [clean] rm -f ./app
task: [build] go build -o app .
task: [run] free | head -n 2
               total        used        free      shared  buff/cache   available
Mem:        65841080    18350600     6361292      129440    41129188    46633780
task: [run] ./app
MemTotal=65841080KB(64297MB), Free=6358804KB(6209MB)
```
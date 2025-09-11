# これは何？

[gopsutil](https://github.com/shirou/gopsutil)を用いてCPU使用率を取得するサンプルです。

```sh
$ task
task: [default] go build -o app .
task: [default] ./app
06:01:48 [4]
06:01:49 [6]
06:01:50 [8]
06:01:51 [6]
06:01:52 [10]
task: [default] ./app -percpu
06:01:54 [4 5]
06:01:55 [5 6]
06:01:56 [6 7]
06:01:57 [8 9]
06:01:58 [2 3]
task: [default] ./app -percpu -spincpu
06:02:00 [83 85]
06:02:01 [84 82]
06:02:02 [85 82]
06:02:03 [84 82]
06:02:04 [85 81]
```

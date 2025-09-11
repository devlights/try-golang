# これは何？

[unix](https://pkg.go.dev/golang.org/x/sys/unix) パッケージの

- Syscall()

を使っているサンプルです。

```sh
$ task
task: [default] rm -f ./app
task: [default] go build -o app .
task: [default] ./app
[syscall] r1=31142, r2=0
[pid    ] syscall=31142, os=31142
```

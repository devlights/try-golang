# これは何？

[unix](https://pkg.go.dev/golang.org/x/sys/unix) パッケージの

- SyscallNoError()

を使っているサンプルです。

ついでに、同じ実装を

- unix.Syscall()
- unixパッケージのラッパー関数 (ex: unix.Listen()など)
- 標準ライブラリの関数（net.Listen()など)

でしてみて、どのように変わるかも試しています。

## 実行結果

```sh
task: [run_noerror] rm -f ./app_noerror
task: [run_noerror] go build -o app_noerror main.go
task: [run_noerror] go build -o app_client client/main.go
task: [run_noerror] ./app_noerror &
task: [run_noerror] sleep 1
task: [run_noerror] ./app_client
[accept] EP: 127.0.0.1:53418
task: [run_noerror] pkill app_noerror
-------------------------------------------------
task: [run_unix] rm -f ./app_unix
task: [run_unix] go build -o app_unix unix/main.go
task: [run_unix] go build -o app_client client/main.go
task: [run_unix] ./app_unix &
task: [run_unix] sleep 1
task: [run_unix] ./app_client
[accept] EP: 127.0.0.1:34092
task: [run_unix] pkill app_unix
-------------------------------------------------
task: [run_stdlib] rm -f ./app_stdlib
task: [run_stdlib] go build -o app_stdlib stdlib/main.go
task: [run_stdlib] go build -o app_client client/main.go
task: [run_stdlib] ./app_stdlib &
task: [run_stdlib] sleep 1
task: [run_stdlib] ./app_client
[accept] EP: 127.0.0.1:34094
task: [run_stdlib] pkill app_stdlib
```

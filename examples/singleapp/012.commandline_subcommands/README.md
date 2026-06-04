# これは何？

```git``` コマンドのように サブコマンド を持つコマンドラインを ```flag``` パッケージを使って構築する方法についてのメモです。

実行すると以下のように表示されます。

```sh
$ task
task: [default] go build -o app
task: [default] ./app
Usage: app <cmd1|cmd2|version|help>
Usage of cmd1:
  -c int
        option c
Usage of cmd2:
  -f string
        option f
task: [default] ./app help
Usage: app <cmd1|cmd2|version|help>
Usage of cmd1:
  -c int
        option c
Usage of cmd2:
  -f string
        option f
task: [default] ./app version
version: vX.Y.Z (fe9c8eb719af1f17bcd9e2e7174c791e7ecd74c0)
task: [default] ./app cmd1 -c 100
option c is 100
task: [default] ./app cmd2 -f helloworld
option f is helloworld
task: [default] ./app cmd1 -f helloworld
flag provided but not defined: -f
Usage of cmd1:
  -c int
        option c
```

## 参考情報

- https://pkg.go.dev/flag@go1.22.1#FlagSet
- https://gobyexample.com/command-line-subcommands

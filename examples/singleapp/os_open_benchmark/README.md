# これは何？

ファイル処理をする際に、[os.Open](https://pkg.go.dev/os@go1.23.5#Open)を毎回実施するのと、開きっぱなしで処理するのとでは、どれくらいパフォーマンスが違うのかについてのサンプルです。

以下の記事を見て、自分用にサンプルとしてここに追加しました。

- [os.Open自体のオーバーヘッドについて](https://qiita.com/Uchijo/items/9337a199040e06b96118)

## 実行結果

```sh
 $ task -d examples/singleapp/os_open_benchmark/
task: [default] go test . -bench=. -run=^$
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/os_open_benchmark
cpu: AMD EPYC 7B13
BenchmarkOsOpenEvery-16           118224              9920 ns/op
BenchmarkOsOpenKeep-16            751408              1526 ns/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/os_open_benchmark    3.862s
```
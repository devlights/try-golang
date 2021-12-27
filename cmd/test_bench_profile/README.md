# これは何？

Go で test, benchmark, profile を行うサンプルです。

## テスト

```sh
$ make test
$ make test_verbose
```

## ベンチマーク

```sh
$ make bench
$ make bench_only
$ make bench_mem
```

## プロファイル

最初にプロファイル結果を出力しておく必要があります。

```sh
$ make profile
```

その後、それぞれのプロファイル結果を見ることが出来ます。

```sh
$ make pprof_cpu
$ make pprof_mem
```

# 参考情報

- [Escape $ dollar sign on Makefiles](https://til.hashrocket.com/posts/k3kjqxtppx-escape-dollar-sign-on-makefiles)
- [Add a test](https://go.dev/doc/tutorial/add-a-test)
- [How to write Go code](https://go.dev/doc/code#Testing)
- [TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Profiling a Go program](https://pkg.go.dev/runtime/pprof#hdr-Profiling_a_Go_program)
- [google/pprof](https://github.com/google/pprof/blob/master/doc/README.md)
- [Profile your golang benchmark with pprof](https://medium.com/@felipedutratine/profile-your-benchmark-with-pprof-fb7070ee1a94)
- [How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
- [pprofでのプロファイル(計測)のやり方を改めて整理した](https://qiita.com/momotaro98/items/bd24a5d4603e378cc357)
- [Goのpprofの使い方【基礎編】](https://christina04.hatenablog.com/entry/golang-pprof-basic)
- [go標準のbenchmark機能の使い方](https://qiita.com/marnie_ms4/items/7014563083ca1d824905)
- [Go pprof 入門編 (CPU Profile とコマンドラインツール)](https://www.klab.com/jp/blog/tech/2015/1047666035.html)
- [Goメモ-138 (staticcheck で警告をコメントで抑制する)](https://devlights.hatenablog.com/entry/2021/03/31/235948)

# これは何？

バイトスライスに対して

- fmt.Sprintf
- fmt.Appendf
- 直接appendで追加していく

のどれが最も速いのかをベンチマークしてみたものです。

[slogのハンドラ作成ガイドドキュメント](https://github.com/golang/example/blob/master/slog-handler-guide/README.md#speed)に記載があったので試してみました。

以下、Gitpod上で試してみた結果です。

```sh
$ task
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/speed_diff_append_string_to_byteslice
cpu: AMD EPYC 7B13
BenchmarkUseFmtSprintf-16        3160230               320.0 ns/op
BenchmarkUseFmtAppendf-16        9861033               133.7 ns/op
BenchmarkUseDirectAppend-16     21009861                83.87 ns/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/speed_diff_append_string_to_byteslice        7.914s
```

処理は冗長になってしまいますが、直接appendが最も速いです。

## SeeAlso

- https://gist.github.com/devlights/ffd22f78297a563c9bebcb9a9baa7f5f

# これは何？

[go:linknameコンパイラディレクティブ](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) を利用して

```time.Now()``` を置き換えるサンプルです。参考情報に上げているサイトの記事がとても勉強になりました。

実行すると以下のようになります。

```sh
$ task
task: [build] go build -o app
task: [default] ./app
2000-01-01 00:00:00 +0000 UTC
```

```time.Now()``` の結果を固定日時になるように置き換えているので、どのタイミングでtime.Now()を呼んでも同じ日時となります。

## 参考情報

- https://www.sobyte.net/post/2022-07/go-linkname/
- https://zenn.dev/sasakiki/articles/a838380540245d
- https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives


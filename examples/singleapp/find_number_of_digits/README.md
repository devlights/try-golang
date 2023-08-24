# これは何？

特定の数値の桁数を求める方法として

- 一度文字列化し、その文字数を数える
- Log10の値をFloor+1する

という２つの方法で試した場合の速度差を見るサンプルです。

Gitpod上で実行すると、例えば以下のような結果が得られます。

```sh
$ task 
[ToString] 9    [Log10] 9
----------------------------------------
=== RUN   TestFn
--- PASS: TestFn (0.00s)
PASS
ok      command-line-arguments  0.002s
----------------------------------------
goos: linux
goarch: amd64
cpu: AMD EPYC 7B13
BenchmarkUseToString-16         34382973                35.82 ns/op
BenchmarkUseLog10-16            96657720                12.26 ns/op
PASS
ok      command-line-arguments  2.473s
```

後者の方が倍以上速いことが分かります。

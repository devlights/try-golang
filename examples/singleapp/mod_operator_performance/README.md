# Overview

Reddit にて、以下のスレッドを発見。

[Go is 2-3 times slower than JS in a similar code. What makes Go slow in this specific code?](https://www.reddit.com/r/golang/comments/11spdom/go_is_23_times_slower_than_js_in_a_similar_code://www.reddit.com/r/golang/comments/11spdom/go_is_23_times_slower_than_js_in_a_similar_code/)

IntelのCPUだと、```%``` を使った計算がとても遅くなるとの話。

内容を見ると、int と uint32 (int32) では、速度が全然異なる状態であると判明している。

どうも、intの状態だと実行時にint64で解釈されてしまうため、遅くなるとのこと。

以下に、AMDとIntelで試してみた結果を記載する。

## CPU: AMD EPYC 7B13

```sh
$ task -d examples/singleapp/mod_operator_performance/
task: [run] go test . -bench . -benchtime 10s
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/mod_operator_performance
cpu: AMD EPYC 7B13
BenchmarkModOperatorInt-16                   252          46435229 ns/op
BenchmarkModOperatorUInt32-16                300          39887931 ns/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/mod_operator_performance     32.641s
```

## CPU: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz

```sh
$ task -d .\examples\singleapp\mod_operator_performance\
task: [run] go test . -bench . -benchtime 10s
goos: windows
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/mod_operator_performance
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkModOperatorInt-8             84         160548945 ns/op
BenchmarkModOperatorUInt32-8         272          42824870 ns/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/mod_operator_performance     30.474s
```

たしかに、Intelの場合、差がとても大きい。


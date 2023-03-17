# Overview

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


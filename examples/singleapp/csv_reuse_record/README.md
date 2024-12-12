# これは何？

```encoding/csv.ReuseRecord``` を有効・無効にした場合のベンチマークです。

```sh
$ task benchmark
task: [benchmark] go test -bench=. -run='^$' -benchmem -count 1 -benchtime 10s
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/csv_reuse_record
cpu: AMD EPYC 7B13
Benchmark_Csv_ReuseRecord-16                  96         117208193 ns/op        39519455 B/op    1991002 allocs/op
Benchmark_Csv_No_ReuseRecord-16               91         136211490 ns/op        69383924 B/op    2115437 allocs/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/csv_reuse_record     23.908s
```

## REFERENCES

- [pkg.go.dev/testing](https://pkg.go.dev/testing#hdr-Benchmarks)
- [Testing flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags)
- [benchstat/go tool traceコマンドをつかったベンチマークの可視化](https://budougumi0617.github.io/2020/12/04/goroutine_tuning_with_benchmark_benchstat_trace/)

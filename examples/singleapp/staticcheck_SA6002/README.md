
```sh
$ task
task: [default] go test -benchmem -run=^$ -bench .
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/staticcheck_SA6002
cpu: AMD EPYC 7B13
BenchmarkStaticCheckSA6002/alloc-16                  549           2159299 ns/op         5241629 B/op         33 allocs/op
BenchmarkStaticCheckSA6002/buffer-16                1678            719659 ns/op               0 B/op          0 allocs/op
BenchmarkStaticCheckSA6002/pool-sa6002-ok-16                1646            781101 ns/op             613 B/op          0 allocs/op
BenchmarkStaticCheckSA6002/pool-sa6002-ng-16                2852            554210 ns/op             378 B/op          1 allocs/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/staticcheck_SA6002   5.699s
```

https://github.com/dominikh/go-tools/issues/1336#issuecomment-1331206290 のコードコメントの日本語訳

```
                // You might be tempted to simplify this by just passing &outBuf to Put,
                // but that would make the local copy of the outBuf slice header escape
                // to the heap, causing an allocation. Instead, we keep around the
                // pointer to the slice header returned by Get, which is already on the
                // heap, and overwrite and return that.
```

```
// ここで &outBuf を Put に渡すことで簡略化したくなるかもしれませんが、
// それでは outBuf スライスヘッダのローカルコピーがヒープに逃げてしまい、
// メモリアロケーションが発生してしまいます。代わりに、Get によって返される
// スライスヘッダへのポインタを保持し、それを上書きして返します。
// そのポインタはすでにヒープ上にあります。
```

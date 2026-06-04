# これは何？

[filepath.Walk](https://pkg.go.dev/path/filepath@go1.22.4#Walk) の説明に以下のように記載されている。

> Walk is less efficient than WalkDir, introduced in Go 1.16, which avoids calling os.Lstat on every visited file or directory.

filepath.Walk が最初から存在しているAPI。[filepath.WalkDir](https://pkg.go.dev/path/filepath@go1.22.4#WalkDir) が Go 1.16 で追加されたももの。

[filepath.Walk](https://pkg.go.dev/path/filepath@go1.22.4#Walk) は、訪問したリソースに対して [os.Lstat](https://pkg.go.dev/os@go1.22.4#Lstat) を呼び出すようになっているため、少し非効率であると記載されている。

実際にベンチマークして算出してみた。

```sh
$ task
task: [default] go test -count 3 -run '^$' -benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/singleapp/walk_walkdir_benchmark
cpu: AMD EPYC 7B13
BenchmarkWalk-16             129           8551681 ns/op          629966 B/op       8957 allocs/op
BenchmarkWalk-16             141           8357755 ns/op          630023 B/op       8957 allocs/op
BenchmarkWalk-16             147           8338617 ns/op          630068 B/op       8957 allocs/op
BenchmarkWalkDir-16          237           4763289 ns/op          348053 B/op       8205 allocs/op
BenchmarkWalkDir-16          248           4924557 ns/op          348213 B/op       8205 allocs/op
BenchmarkWalkDir-16          229           5093532 ns/op          348304 B/op       8205 allocs/op
PASS
ok      github.com/devlights/try-golang/examples/singleapp/walk_walkdir_benchmark       11.160s
```

確かに [filepath.WalkDir](https://pkg.go.dev/path/filepath@go1.22.4#WalkDir) の方が処理速度が速く、メモリ割り当ても効率的である。

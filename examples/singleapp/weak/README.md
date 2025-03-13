# Go 1.24 で追加された weak パッケージのサンプル

Go 1.24 にて [weak](https://pkg.go.dev/weak@go1.24.1) パッケージが追加された。弱参照をサポートするライブラリ。

```sh
$ task
task: [default] go run main.go
[init  ] HeapAlloc:    185.22 KB (ヒープメモリ)
[before] HeapAlloc:     32.19 MB (ヒープメモリ)
[after ] HeapAlloc:     32.19 MB (ヒープメモリ)
object is nil? ==> false
task: [default] sleep 1
task: [default] go run main.go -weakref
[init  ] HeapAlloc:    185.50 KB (ヒープメモリ)
[before] HeapAlloc:     32.19 MB (ヒープメモリ)
[after ] HeapAlloc:    194.81 KB (ヒープメモリ)
object is nil? ==> true
```

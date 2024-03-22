# これは何？

```expvar``` パッケージのサンプルです。

実行すると以下のようになります。

```sh
$ task run
task: [build] go build
task: [run] ./expvar opt1 opt2 &
task: [run] for i in {1..3} ; do curl -sS http://localhost:8888/debug/vars | jq 'del(.memstats)'; sleep 2; done
{
  "cmdline": [
    "./expvar",
    "opt1",
    "opt2"
  ],
  "counter": 0,
  "message": "",
  "values": {}
}
{
  "cmdline": [
    "./expvar",
    "opt1",
    "opt2"
  ],
  "counter": 2,
  "message": "hello-02",
  "values": {
    "counter": 4,
    "message": "HELLO-02"
  }
}
{
  "cmdline": [
    "./expvar",
    "opt1",
    "opt2"
  ],
  "counter": 4,
  "message": "hello-04",
  "values": {
    "counter": 16,
    "message": "HELLO-04"
  }
}
task: [run] pgrep expvar | xargs kill
```

## 参考情報

- https://pkg.go.dev/expvar@go1.22.1
- https://qiita.com/methane/items/8f56f663d6da4dee9f64

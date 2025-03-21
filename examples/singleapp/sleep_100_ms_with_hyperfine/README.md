# これは何？

```time.Sleep()```を利用して例えば100ms待機と実行した場合

現実的にはどの程度の誤差が出るのかを[hyperfine](https://github.com/sharkdp/hyperfine)を利用して確認するサンプルです。

## 例

以下はGitpod環境で実行した場合

```sh
task: [build] go build -o app
task: [run] hyperfine './app -d 100'
Benchmark 1: ./app -d 100
  Time (mean ± σ):     102.2 ms ±   0.2 ms    [User: 1.5 ms, System: 0.9 ms]
  Range (min … max):   101.9 ms … 103.2 ms    29 runs
 
task: [run] hyperfine './app -d 99'
Benchmark 1: ./app -d 99
  Time (mean ± σ):     101.1 ms ±   0.2 ms    [User: 1.5 ms, System: 1.0 ms]
  Range (min … max):   100.8 ms … 101.5 ms    29 runs
 
task: [run] hyperfine './app -d 98'
Benchmark 1: ./app -d 98
  Time (mean ± σ):     100.1 ms ±   0.1 ms    [User: 1.6 ms, System: 0.9 ms]
  Range (min … max):    99.8 ms … 100.3 ms    29 runs
 
task: [run] hyperfine './app -d 97'
Benchmark 1: ./app -d 97
  Time (mean ± σ):      99.1 ms ±   0.1 ms    [User: 1.6 ms, System: 0.8 ms]
  Range (min … max):    98.8 ms …  99.3 ms    29 runs
 
task: [run] hyperfine './app -d 96'
Benchmark 1: ./app -d 96
  Time (mean ± σ):      98.0 ms ±   0.2 ms    [User: 1.5 ms, System: 0.8 ms]
  Range (min … max):    97.8 ms …  98.5 ms    30 runs
 
task: [run] hyperfine './app -d 95'
Benchmark 1: ./app -d 95
  Time (mean ± σ):      97.1 ms ±   0.2 ms    [User: 1.6 ms, System: 0.8 ms]
  Range (min … max):    96.8 ms …  97.5 ms    30 runs
 
task: [run] hyperfine 'sleep 0.1'
Benchmark 1: sleep 0.1
  Time (mean ± σ):     101.3 ms ±   0.1 ms    [User: 0.8 ms, System: 0.6 ms]
  Range (min … max):   101.1 ms … 101.7 ms    29 runs
```

もっとも100msに近いのは ```time.Sleep(98*time.Millisecond)``` となった。

また、以下はWindows上のWSL（Ubuntu Linux 24.04 LTS) での結果。

```sh
task: [build] go build
task: [run] hyperfine './app -d 100'
Benchmark 1: ./app -d 100
  Time (mean ± σ):     104.7 ms ±   0.8 ms    [User: 4.6 ms, System: 0.4 ms]
  Range (min … max):   103.6 ms … 106.5 ms    28 runs

task: [run] hyperfine './app -d 99'
Benchmark 1: ./app -d 99
  Time (mean ± σ):     103.8 ms ±   0.9 ms    [User: 4.1 ms, System: 0.8 ms]
  Range (min … max):   101.5 ms … 105.8 ms    29 runs

task: [run] hyperfine './app -d 98'
Benchmark 1: ./app -d 98
  Time (mean ± σ):     102.7 ms ±   0.8 ms    [User: 3.7 ms, System: 0.9 ms]
  Range (min … max):   101.4 ms … 104.6 ms    29 runs

task: [run] hyperfine './app -d 97'
Benchmark 1: ./app -d 97
  Time (mean ± σ):     101.9 ms ±   1.0 ms    [User: 3.8 ms, System: 0.8 ms]
  Range (min … max):    99.7 ms … 103.5 ms    29 runs

task: [run] hyperfine './app -d 96'
Benchmark 1: ./app -d 96
  Time (mean ± σ):     101.8 ms ±   1.0 ms    [User: 4.5 ms, System: 1.1 ms]
  Range (min … max):   100.0 ms … 103.8 ms    29 runs

task: [run] hyperfine './app -d 95'
Benchmark 1: ./app -d 95
  Time (mean ± σ):     100.8 ms ±   0.7 ms    [User: 4.4 ms, System: 1.3 ms]
  Range (min … max):    98.4 ms … 102.0 ms    30 runs

task: [run] hyperfine 'sleep 0.1'
Benchmark 1: sleep 0.1
  Time (mean ± σ):     104.0 ms ±   0.6 ms    [User: 3.4 ms, System: 0.4 ms]
  Range (min … max):   102.3 ms … 105.0 ms    29 runs
```

こちらは95ミリ待機がもっとも100ミリ秒に近い結果となった。

このように実働する環境によって、誤差がでるものなので、実際にアプリケーションが稼働している環境でのベンチマーク作業はとても大事。


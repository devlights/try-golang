# syscall.Nanosleep

time.Sleep()で問題ない場合が９割９部であるが、極稀に極めて正確にタイミングを制御しないといけない場合、システムコールを直接呼び出す方が極僅かであるが正確性が上がるときがある。

time.Sleep()は内部でGoのスケジューラと強調して動作しているため、ほんの僅かのオーバーヘッドがある。

syscall.Nanosleep()は、名前の通り低レベルのシステムコールを直接呼び出すのと同じことになる。

## 5ミリ秒での計測

計測時間が極端に短いので、hyperfineにてマイクロ秒の単位で出力するようオプション制御。

```sh
$ task
task: [build] go build -o app .
task: [run] hyperfine --time-unit microsecond --warmup 3 --runs 10 './app -syscall -val 5'
Benchmark 1: ./app -syscall -val 5
  Time (mean ± σ):     6598.1 µs ±  76.6 µs    [User: 1392.5 µs, System: 643.6 µs]
  Range (min … max):   6538.4 µs … 6733.6 µs    10 runs
 
-----------------------
task: [run] hyperfine --time-unit microsecond --warmup 3 --runs 10 './app -val 5'
Benchmark 1: ./app -val 5
  Time (mean ± σ):     6749.0 µs ± 142.1 µs    [User: 1176.5 µs, System: 829.2 µs]
  Range (min … max):   6530.1 µs … 6944.2 µs    10 runs
```

1. **実行時間の比較**:
   - `syscall.Nanosleep`: 平均6,598.1µs（約6.6ms）
   - `time.Sleep`: 平均6,749.0µs（約6.7ms）
   - 差: 約150.9µs（`syscall.Nanosleep`がわずかに速い）

2. **指定時間との差異**:
   - 指定値: 5,000µs（5ms）
   - `syscall.Nanosleep`: +1,598.1µs（31.96%長い）
   - `time.Sleep`: +1,749.0µs（34.98%長い）

3. **安定性の比較**:
   - `syscall.Nanosleep`の標準偏差: 76.6µs
   - `time.Sleep`の標準偏差: 142.1µs
   - `syscall.Nanosleep`の方が約1.86倍安定している

4. **CPU使用状況**:
   - `syscall.Nanosleep`: User 1,392.5µs, System 643.6µs（計2,036.1µs）
   - `time.Sleep`: User 1,176.5µs, System 829.2µs（計2,005.7µs）
   - ユーザー時間は`syscall.Nanosleep`の方が長いが、システム時間は`time.Sleep`の方が長い

5. **最小・最大値の範囲**:
   - `syscall.Nanosleep`: 6,538.4µs～6,733.6µs（範囲: 195.2µs）
   - `time.Sleep`: 6,530.1µs～6,944.2µs（範囲: 414.1µs）
   - `time.Sleep`の変動幅が2.12倍大きい


## 10ミリ秒で計測

計測には [hyperfine](https://github.com/sharkdp/hyperfine) を使用。

```sh
$ task run
task: [build] go build -o app .
task: [run] hyperfine --warmup 3 --runs 10 './app -syscall'
Benchmark 1: ./app -syscall
  Time (mean ± σ):      11.6 ms ±   0.1 ms    [User: 1.9 ms, System: 0.4 ms]
  Range (min … max):    11.5 ms …  11.7 ms    10 runs
 
-----------------------
task: [run] hyperfine --warmup 3 --runs 10 './app'
Benchmark 1: ./app
  Time (mean ± σ):      13.4 ms ±   3.4 ms    [User: 1.7 ms, System: 1.1 ms]
  Range (min … max):    11.8 ms …  23.1 ms    10 runs
```

## 30ミリ秒で計測

```sh
$ task
task: [build] go build -o app .
task: [run] hyperfine --warmup 3 --runs 10 './app -syscall'
Benchmark 1: ./app -syscall
  Time (mean ± σ):      31.8 ms ±   0.2 ms    [User: 1.6 ms, System: 0.6 ms]
  Range (min … max):    31.6 ms …  32.1 ms    10 runs
 
-----------------------
task: [run] hyperfine --warmup 3 --runs 10 './app'
Benchmark 1: ./app
  Time (mean ± σ):      31.9 ms ±   0.2 ms    [User: 1.6 ms, System: 0.7 ms]
  Range (min … max):    31.7 ms …  32.3 ms    10 runs
```


## 50ミリ秒で計測

```sh
$ task
task: [build] go build -o app .
task: [run] hyperfine --warmup 3 --runs 10 './app -syscall -val 50'
Benchmark 1: ./app -syscall -val 50
  Time (mean ± σ):      52.1 ms ±   0.3 ms    [User: 1.6 ms, System: 0.9 ms]
  Range (min … max):    51.8 ms …  52.5 ms    10 runs
 
-----------------------
task: [run] hyperfine --warmup 3 --runs 10 './app -val 50'
Benchmark 1: ./app -val 50
  Time (mean ± σ):      51.9 ms ±   0.1 ms    [User: 1.5 ms, System: 0.8 ms]
  Range (min … max):    51.8 ms …  52.0 ms    10 runs
```

## 100ミリ秒で計測

```sh
$ task
task: [build] go build -o app .
task: [run] hyperfine --warmup 3 --runs 10 './app -syscall -val 100'
Benchmark 1: ./app -syscall -val 100
  Time (mean ± σ):     101.5 ms ±   0.2 ms    [User: 1.4 ms, System: 1.2 ms]
  Range (min … max):   101.2 ms … 101.7 ms    10 runs
 
-----------------------
task: [run] hyperfine --warmup 3 --runs 10 './app -val 100'
Benchmark 1: ./app -val 100
  Time (mean ± σ):     102.1 ms ±   0.1 ms    [User: 1.6 ms, System: 0.9 ms]
  Range (min … max):   102.0 ms … 102.4 ms    10 runs
```

## 500ミリ秒で計測

```sh
$ task
task: [build] go build -o app .
task: [run] hyperfine --warmup 3 --runs 10 './app -syscall -val 500'
Benchmark 1: ./app -syscall -val 500
  Time (mean ± σ):     502.6 ms ±   0.4 ms    [User: 1.7 ms, System: 1.3 ms]
  Range (min … max):   502.0 ms … 503.5 ms    10 runs
 
-----------------------
task: [run] hyperfine --warmup 3 --runs 10 './app -val 500'
Benchmark 1: ./app -val 500
  Time (mean ± σ):     502.6 ms ±   0.4 ms    [User: 1.3 ms, System: 1.3 ms]
  Range (min … max):   502.0 ms … 503.5 ms    10 runs
```

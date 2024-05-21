package metricsop

import (
	"context"
	"math"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/devlights/gomy/output"
)

// Cpu は、runtime/metrics を利用してCPU関連の情報を取得するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime/metrics@latest
func Cpu() error {
	var (
		items = []string{
			"/cgo/go-to-c-calls:calls",                   // 現在のプロセスによるGoからCへの呼び出し回数
			"/cpu/classes/gc/mark/assist:cpu-seconds",    // ゴルーチンがGCを実行するのに費やしたCPU時間の見積もり
			"/cpu/classes/gc/mark/dedicated:cpu-seconds", // GCタスクを実行するために費やされたCPU時間の合計の見積もり
			"/cpu/classes/gc/mark/idle:cpu-seconds",      // GCタスクの実行に費やされたCPU時間の合計の見積もり
			"/cpu/classes/gc/pause:cpu-seconds",          // GCによって一時停止されたアプリケーションに費やされた推定総CPU時間
			"/cpu/classes/gc/total:cpu-seconds",          // GCタスクの実行に費やされたCPU時間の見積もり
			"/cpu/classes/idle:cpu-seconds",              // GoまたはGoランタイムコードの実行に使用されないCPU時間の推定値
			"/cpu/classes/total:cpu-seconds",             // ユーザーGoコードまたはGoランタイムに使用可能なCPU時間の合計
			"/cpu/classes/user:cpu-seconds",              // ユーザーGoコードの実行に費やされたCPU時間の合計の見積もり
		}
		samples = make([]metrics.Sample, len(items))
		bigdata = make([]byte, 1<<28)
	)

	for i, name := range items {
		samples[i].Name = name
	}

	var (
		ctx, cxl = context.WithTimeout(context.Background(), 1*time.Second)
		ready    = make(chan bool)
		busyfn   = func(ctx context.Context, ready chan<- bool) {
			var i uint64

			ready <- true
			for {
				select {
				case <-ctx.Done():
					return
				default:
					switch i {
					case math.MaxUint64:
						i = 0
					default:
						i += uint64(1)
					}
				}
			}
		}
	)
	defer cxl()

	go busyfn(ctx, ready)
	<-ready

	runtime.GC()
	metrics.Read(samples)

	for _, s := range samples {
		output.Stdoutl("[Name ]", s.Name)
		output.Stdoutf("[Value]", "%+v\n", s.Value)
		output.StdoutHr()
	}

	output.Stdoutl("[Buffer]", len(bigdata))

	<-ctx.Done()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: metrics_cpu

	   [Name] "metrics_cpu"
	   [Name ]              /cgo/go-to-c-calls:calls
	   [Value]              {kind:1 scalar:1 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/gc/mark/assist:cpu-seconds
	   [Value]              {kind:2 scalar:4551153223746165794 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/gc/mark/dedicated:cpu-seconds
	   [Value]              {kind:2 scalar:4562351584908057236 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/gc/mark/idle:cpu-seconds
	   [Value]              {kind:2 scalar:0 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/gc/pause:cpu-seconds
	   [Value]              {kind:2 scalar:4566896090190411183 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/gc/total:cpu-seconds
	   [Value]              {kind:2 scalar:4569689296178052284 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/idle:cpu-seconds
	   [Value]              {kind:2 scalar:4630569171334083814 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/total:cpu-seconds
	   [Value]              {kind:2 scalar:4630572714109167296 pointer:<nil>}
	   --------------------------------------------------
	   [Name ]              /cpu/classes/user:cpu-seconds
	   [Value]              {kind:2 scalar:4581969798399309374 pointer:<nil>}
	   --------------------------------------------------
	   [Buffer]             268435456


	   [Elapsed] 1.002128694s
	*/

}

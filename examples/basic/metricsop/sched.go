package metricsop

import (
	"context"
	"runtime"
	"runtime/metrics"
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// Sched は、runtime/metrics を利用してスケジューラ関連の情報を取得するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime/metrics@latest
func Sched() error {
	var (
		items = []string{
			"/sched/gomaxprocs:threads",      // ユーザーレベルのGoコードを同時に実行できるオペレーティング・システムのスレッド数
			"/sched/goroutines:goroutines",   // 生きているゴルーチンの数
			"/sched/latencies:seconds",       // ゴルーチンが実際に実行される前に、スケジューラ内で実行可能な状態で過ごした時間の分布
			"/sync/mutex/wait/total:seconds", // ゴルーチンがsync.Mutex、sync.RWMutex、またはランタイム内部ロックでブロックされた時間の累計
		}
		samples = make([]metrics.Sample, len(items))
	)

	for i, name := range items {
		samples[i].Name = name
	}

	var (
		ctx, cxl = context.WithTimeout(context.Background(), 2*time.Second)
		ready    = make(chan bool)
		lock     sync.Mutex
		busyfn   = func(ctx context.Context, ready <-chan bool) {
			<-ready

			lock.Lock()
			time.Sleep(100 * time.Millisecond)
			lock.Unlock()

			<-ctx.Done()
		}
	)
	defer cxl()

	for range runtime.GOMAXPROCS(0) - 1 {
		go busyfn(ctx, ready)
	}
	close(ready)

	<-time.After(1 * time.Second)

	metrics.Read(samples)
	for _, s := range samples {
		output.Stdoutl("[Name ]", s.Name)

		switch s.Value.Kind() {
		case metrics.KindUint64:
			output.Stdoutf("[Value]", "%v\n", s.Value.Uint64())
		case metrics.KindFloat64:
			output.Stdoutf("[Value]", "%v\n", s.Value.Float64())
		case metrics.KindFloat64Histogram:
			output.Stdoutf("[Value]", "Bucket Count: %d\n", len(s.Value.Float64Histogram().Buckets)-2)
		default:
			output.Stdoutl("[Value]", "INVALID")
		}

		output.StdoutHr()
	}

	<-ctx.Done()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: metrics_sched

	   [Name] "metrics_sched"
	   [Name ]              /sched/gomaxprocs:threads
	   [Value]              16
	   --------------------------------------------------
	   [Name ]              /sched/goroutines:goroutines
	   [Value]              16
	   --------------------------------------------------
	   [Name ]              /sched/latencies:seconds
	   [Value]              Bucket Count: 161
	   --------------------------------------------------
	   [Name ]              /sync/mutex/wait/total:seconds
	   [Value]              4.010605616
	   --------------------------------------------------


	   [Elapsed] 2.000425097s
	*/

}

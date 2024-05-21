package metricsop

import (
	"runtime"
	"runtime/metrics"

	"github.com/devlights/gomy/output"
)

// Heap は、runtime/metrics を利用してヒープメモリ関連の情報を取得するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime/metrics@latest
func Heap() error {
	var (
		items = []string{
			"/memory/classes/heap/free:bytes",     // 完全に空いていて、システムに戻す資格があるが、戻されていないメモリ
			"/memory/classes/heap/objects:bytes",  // ガベージコレクタによってまだ解放されていないメモリ
			"/memory/classes/heap/released:bytes", // 完全に解放され、システムに戻されたメモリ
			"/memory/classes/heap/stacks:bytes",   // スタック領域として予約されている、ヒープから割り当てられたメモリ、現在使用中であるか否かを問わない
			"/memory/classes/heap/unused:bytes",   // ヒープオブジェクトのために予約されているが、現在ヒープオブジェクトを保持するために使われていないメモリ
			"/gc/heap/allocs:bytes",               // アプリケーションによってヒープに割り当てられたメモリの累計。
			"/gc/heap/allocs:objects",             // アプリケーションによって引き起こされたヒープ割り当ての累積カウント
			"/gc/heap/frees:bytes",                // ガベージコレクタによって解放されたヒープメモリの累計
			"/gc/heap/frees:objects",              // ストレージがガベージコレクタによって解放されたヒープ割り当ての累積カウント
			"/gc/heap/goal:bytes",                 // GCサイクル終了時のヒープサイズ目標
			"/gc/heap/live:bytes",                 // 前回のGCでマークされたライブオブジェクトが占有するヒープメモリ
			"/gc/heap/objects:objects",            // ヒープメモリを占有しているオブジェクトの数
		}
		samples = make([]metrics.Sample, len(items))
		bigdata = make([]byte, 1<<28)
	)

	for i, name := range items {
		samples[i].Name = name
	}

	runtime.GC()
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

	output.Stdoutl("[Buffer]", len(bigdata))

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: metrics_heap

		[Name] "metrics_heap"
		[Name ]              /memory/classes/heap/free:bytes
		[Value]              237232128
		--------------------------------------------------
		[Name ]              /memory/classes/heap/objects:bytes
		[Value]              395616
		--------------------------------------------------
		[Name ]              /memory/classes/heap/released:bytes
		[Value]              33882112
		--------------------------------------------------
		[Name ]              /memory/classes/heap/stacks:bytes
		[Value]              425984
		--------------------------------------------------
		[Name ]              /memory/classes/heap/unused:bytes
		[Value]              693920
		--------------------------------------------------
		[Name ]              /gc/heap/allocs:bytes
		[Value]              268953040
		--------------------------------------------------
		[Name ]              /gc/heap/allocs:objects
		[Value]              1668
		--------------------------------------------------
		[Name ]              /gc/heap/frees:bytes
		[Value]              268557424
		--------------------------------------------------
		[Name ]              /gc/heap/frees:objects
		[Value]              459
		--------------------------------------------------
		[Name ]              /gc/heap/goal:bytes
		[Value]              4194304
		--------------------------------------------------
		[Name ]              /gc/heap/live:bytes
		[Value]              395904
		--------------------------------------------------
		[Name ]              /gc/heap/objects:objects
		[Value]              1209
		--------------------------------------------------
		[Buffer]             268435456


		[Elapsed] 4.183119ms
	*/
}

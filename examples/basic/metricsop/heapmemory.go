package metricsop

import (
	"runtime"
	"runtime/metrics"

	"github.com/devlights/gomy/output"
)

// HeapMemory は、runtime/metrics を利用してヒープメモリ関連の情報を取得するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime/metrics@latest
func HeapMemory() error {
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
		output.Stdoutf("[Value]", "%+v\n", s.Value)
		output.StdoutHr()
	}

	output.Stdoutl("[Buffer]", len(bigdata))

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: metrics_heapmemory

		[Name] "metrics_heapmemory"
		[Name ]              /memory/classes/heap/free:bytes
		[Value]              {kind:1 scalar:268656640 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /memory/classes/heap/objects:bytes
		[Value]              {kind:1 scalar:398432 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /memory/classes/heap/released:bytes
		[Value]              {kind:1 scalar:2433024 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /memory/classes/heap/stacks:bytes
		[Value]              {kind:1 scalar:458752 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /memory/classes/heap/unused:bytes
		[Value]              {kind:1 scalar:682912 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/allocs:bytes
		[Value]              {kind:1 scalar:268953784 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/allocs:objects
		[Value]              {kind:1 scalar:1662 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/frees:bytes
		[Value]              {kind:1 scalar:268555352 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/frees:objects
		[Value]              {kind:1 scalar:450 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/goal:bytes
		[Value]              {kind:1 scalar:4194304 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/live:bytes
		[Value]              {kind:1 scalar:398720 pointer:<nil>}
		--------------------------------------------------
		[Name ]              /gc/heap/objects:objects
		[Value]              {kind:1 scalar:1212 pointer:<nil>}
		--------------------------------------------------
		[Buffer]             268435456


		[Elapsed] 2.96696ms
	*/
}

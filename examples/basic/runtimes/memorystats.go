package runtimes

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// RuntimeMemoryStats は、runtime.MemoryStats() のサンプルです.
//
// REFERENCES::
//   - https://golangcode.com/print-the-current-memory-usage/
func RuntimeMemoryStats() error {
	var (
		rootCtx         = context.Background()
		mainCtx, cancel = context.WithCancel(rootCtx)
		wg              = &sync.WaitGroup{}
	)

	// 初期の状態を表示
	runtime.GC()
	printMemoryStats("init")

	// ---------------------------------------
	// データを500ミリ秒毎に増やしていく処理
	// ---------------------------------------
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		var (
			tick  = time.Tick(500 * time.Millisecond) //lint:ignore SA1015 time.Tickはリークの危険があることは認識済み
			count = 0
			items = make([][]byte, 0, 5)
		)

		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick:
				count++
				data := make([]byte, 0, 1024*1024)
				//lint:ignore SA4010 意味がないことは承知済み
				items = append(items, data)

				// output.Stderrf("[append]", "count=%d\ttick=%v\n", count, t)
			}
		}
	}(mainCtx, wg)

	// ---------------------------------------
	// 現在のメモリ量を2000ミリ秒毎に出力する処理
	// ---------------------------------------
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		var (
			tick  = time.Tick(2000 * time.Millisecond) //lint:ignore SA1015 time.Tickはリークの危険があることは認識済み
			count = 0
		)

		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick:
				count++
				printMemoryStats(fmt.Sprintf("(%d)", count))
			}
		}
	}(mainCtx, wg)

	// 10秒したら終わり
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()

	// 最後の状態を表示
	printMemoryStats("latest")

	// GC後の状態を表示
	runtime.GC()
	printMemoryStats("after runtime.GC()")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_memorystats

	   [Name] "runtime_memorystats"
	   init                 ----------------------------
	   Alloc                365
	   HeapAlloc            365
	   TotalAlloc           437
	   HeapObjects          1
	   Sys                  9253
	   NumGC                1
	   (1)                  ----------------------------
	   Alloc                4462
	   HeapAlloc            4462
	   TotalAlloc           4538
	   HeapObjects          1
	   Sys                  13413
	   NumGC                2
	   (2)                  ----------------------------
	   Alloc                8559
	   HeapAlloc            8559
	   TotalAlloc           8635
	   HeapObjects          1
	   Sys                  17765
	   NumGC                2
	   (3)                  ----------------------------
	   Alloc                12657
	   HeapAlloc            12657
	   TotalAlloc           12737
	   HeapObjects          1
	   Sys                  22117
	   NumGC                3
	   (4)                  ----------------------------
	   Alloc                16754
	   HeapAlloc            16754
	   TotalAlloc           16834
	   HeapObjects          1
	   Sys                  26213
	   NumGC                3
	   (5)                  ----------------------------
	   Alloc                20851
	   HeapAlloc            20851
	   TotalAlloc           20935
	   HeapObjects          1
	   Sys                  30565
	   NumGC                4
	   latest               ----------------------------
	   Alloc                20852
	   HeapAlloc            20852
	   TotalAlloc           20935
	   HeapObjects          1
	   Sys                  30565
	   NumGC                4
	   after runtime.GC()   ----------------------------
	   Alloc                372
	   HeapAlloc            372
	   TotalAlloc           20940
	   HeapObjects          1
	   Sys                  30565
	   NumGC                5


	   [Elapsed] 10.010817629s
	*/

}

func printMemoryStats(prefix string) {
	// --------------------------------------------------------
	// runtime.MemoryStats() から、現在の割当メモリ量などが取得できる.
	//
	// まず、データの受け皿となる runtime.MemStats を初期化し
	// runtime.ReadMemStats(*runtime.MemStats) を呼び出して
	// 取得する.
	// --------------------------------------------------------
	var (
		ms runtime.MemStats
	)

	output.Stdoutl(prefix, "----------------------------")
	runtime.ReadMemStats(&ms)

	// Alloc は、現在ヒープに割り当てられているメモリ
	// HeapAlloc と同じ.
	output.Stdoutl("Alloc", toKb(ms.Alloc))
	output.Stdoutl("HeapAlloc", toKb(ms.HeapAlloc))

	// TotalAlloc は、ヒープに割り当てられたメモリ量の累積
	// Allocと違い、こちらは増えていくが減ることはない
	output.Stdoutl("TotalAlloc", toKb(ms.TotalAlloc))

	// HeapObjects は、ヒープに割り当てられているオブジェクトの数
	output.Stdoutl("HeapObjects", toKb(ms.HeapObjects))

	// Sys は、OSから割り当てられたメモリの合計量
	output.Stdoutl("Sys", toKb(ms.Sys))

	// NumGC は、実施されたGCの回数
	output.Stdoutl("NumGC", ms.NumGC)
}

func toKb(bytes uint64) uint64 {
	return bytes / 1024
}

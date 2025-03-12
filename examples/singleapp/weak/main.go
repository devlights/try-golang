package main

import (
	"fmt"
	"runtime"
	"weak"
)

// メモリ統計情報を見やすく表示する構造体
type MemStats struct {
	// ヒープに割り当てられたメモリの合計 (bytes)
	HeapAlloc uint64
	// OSから取得したヒープメモリの合計 (bytes)
	HeapSys uint64
	// 現在使用中のヒープオブジェクト数
	HeapObjects uint64
	// ヒープに割り当てられたが解放されていないメモリ (bytes)
	HeapInuse uint64
	// ヒープに割り当てられた未使用メモリ (bytes)
	HeapIdle uint64
	// 次のGCが発生するヒープサイズ (bytes)
	NextGC uint64
	// 最後のGC以降に割り当てられた累積メモリ (bytes)
	TotalAlloc uint64
}

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	const oneMib = 1 << 20
	{
		printMem("init  ", getMemStats())

		x := make([]byte, oneMib<<5)
		strongRef := struct {
			b *[]byte
		}{&x}

		printMem("before", getMemStats())
		runtime.GC()
		printMem("after ", getMemStats())
		runtime.KeepAlive(&strongRef)
	}

	runtime.GC()
	fmt.Println("------------------------------------")

	{
		printMem("init  ", getMemStats())

		x := make([]byte, oneMib<<5)
		weakRef := weak.Make(&x)

		printMem("before", getMemStats())
		runtime.GC()
		printMem("after ", getMemStats())
		runtime.KeepAlive(&weakRef)

		fmt.Printf("object is nil? ==> %v\n", weakRef.Value() == nil)
	}

	return nil
}

// メモリサイズをより読みやすい形式に変換
func formatBytes(bytes uint64) string {
	switch {
	case bytes >= uint64(GB):
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= uint64(MB):
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= uint64(KB):
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}

func printMem(prefix string, stats MemStats) {
	fmt.Printf("[%s] HeapAlloc:   %s (現在ヒープに割り当てられているメモリ)\n", prefix, formatBytes(stats.HeapAlloc))
}

func getMemStats() MemStats {
	var rtm runtime.MemStats

	runtime.ReadMemStats(&rtm)

	return MemStats{
		HeapAlloc:   rtm.HeapAlloc,
		HeapSys:     rtm.HeapSys,
		HeapObjects: rtm.HeapObjects,
		HeapInuse:   rtm.HeapInuse,
		HeapIdle:    rtm.HeapIdle,
		NextGC:      rtm.NextGC,
		TotalAlloc:  rtm.TotalAlloc,
	}
}

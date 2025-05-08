package main

import (
	"log"

	"github.com/prometheus/procfs"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		fs  procfs.FS
		err error
	)
	fs, err = procfs.NewDefaultFS() // デフォルトは /proc を見ている
	if err != nil {
		return err
	}

	var (
		mem procfs.Meminfo
	)
	mem, err = fs.Meminfo()
	if err != nil {
		return err
	}

	var (
		memTotal       = *mem.MemTotal      // システムに搭載されている物理メモリ（RAM）の総量をkB単位
		memTotalBytes  = *mem.MemTotalBytes // システムに搭載されている物理メモリ（RAM）の総量をバイト単位
		freeTotal      = *mem.MemFree       // 空き容量をKB表示
		freeTotalBytes = *mem.MemFreeBytes  // 空き容量をバイト単位
		toMB           = func(v uint64) uint64 { return v >> 20 }
	)
	log.Printf("MemTotal=%dKB(%dMB), Free=%dKB(%dMB)", memTotal, toMB(memTotalBytes), freeTotal, toMB(freeTotalBytes))

	return nil
}

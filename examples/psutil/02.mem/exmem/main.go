//go:build linux

package main

import (
	"log"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	log.SetFlags(log.Ltime)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		vmem   *mem.VirtualMemoryStat
		exVmem *mem.ExVirtualMemory
		err    error
	)
	// マルチプラットフォームで利用出来る方法
	vmem, err = mem.VirtualMemory()
	if err != nil {
		return err
	}

	// 各OSごとに特化している情報を取得する方法
	exVmem, err = mem.NewExLinux().VirtualMemory()
	if err != nil {
		return err
	}

	log.Printf("[NORMAL] %s", vmem)
	log.Printf("[EX    ] %s", exVmem)

	return nil
}

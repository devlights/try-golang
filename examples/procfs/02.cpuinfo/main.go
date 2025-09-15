//go:build linux

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
		cpuInfos []procfs.CPUInfo
	)
	cpuInfos, err = fs.CPUInfo()
	if err != nil {
		return err
	}

	log.Printf("CPU-COUNT: %d", len(cpuInfos))
	for i, cpu := range cpuInfos {
		log.Printf("\t[CPU-%02d] CoreId=%s ModelName=%q, Processor=%02d", i+1, cpu.CoreID, cpu.ModelName, cpu.Processor)
	}

	return nil
}

package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	log.SetFlags(log.Ltime)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	const (
		COUNT = 5
	)
	var (
		interval = 1 * time.Second
		mb       = func(v uint64) uint64 {
			// 概算値で良いので整数除算
			return v >> 20
		}
		done   = make(chan struct{})
		useMem = func(done <-chan struct{}, ready chan<- struct{}, bufSize uint64) {
			buf := make([]byte, bufSize)
			for i := range len(buf) {
				buf[i] = 0
			}

			close(ready)
			<-done
		}
	)
	for range COUNT {
		// 現在のメモリ量を取得
		var (
			vms *mem.VirtualMemoryStat
			err error
		)
		vms, err = mem.VirtualMemory()
		if err != nil {
			return err
		}

		log.Printf("Total=%4dMB\tUsed=%4dMB\tFree=%4dMB", mb(vms.Total), mb(vms.Used), mb(vms.Free))

		// 意図的にメモリを占有していく
		var (
			ready = make(chan struct{})
		)
		go useMem(done, ready, 150<<20)
		<-ready

		time.Sleep(interval)
	}

	close(done)
	time.Sleep(500 * time.Millisecond)

	return nil
}

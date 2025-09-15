//go:build linux

package main

import (
	"flag"
	"syscall"
	"time"
)

type (
	Args struct {
		UseSyscall bool
		Val        int
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.UseSyscall, "syscall", false, "use syscall.nanosleep")
	flag.IntVar(&args.Val, "val", 1, "value")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	switch args.UseSyscall {
	case true:
		//
		// syscall.Nanosleep() を使用
		//
		ts := syscall.NsecToTimespec(int64(args.Val) * time.Millisecond.Nanoseconds())
		_ = syscall.Nanosleep(&ts, nil)
	default:
		//
		// time.Sleep() を使用
		//
		d := time.Duration(args.Val) * time.Millisecond
		time.Sleep(d)
	}

	return nil
}

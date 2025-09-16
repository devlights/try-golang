//go:build unix

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// unix.Syscall() のサンプル
	//
	// unix.Syscall(), unix.Syscall6() は、指定したシステムコールを呼び出すための関数。
	//
	// unix.Syscall() は引数が３つ指定できるシステムコール呼び出し関数。
	// これより多い場合は６つ指定できる unix.Syscall6() を利用する。
	//

	//
	// getpid(2)
	// getpid(2)は引数が無い関数 (pid_t getpid(void);)
	//
	const (
		ZERO uintptr = 0 // unix.Syscall()で指定する際に利用する不要な引数値
	)
	var (
		trap   uintptr = unix.SYS_GETPID
		r1, r2 uintptr
		err    unix.Errno // unix.Syscall()の場合は error ではなく unix.Errno であることに注意
	)
	r1, r2, err = unix.Syscall(trap, ZERO, ZERO, ZERO)
	if err != 0 {
		return fmt.Errorf("getpid: %d", err)
	}

	log.Printf("[syscall] r1=%d, r2=%d", r1, r2)
	log.Printf("[pid    ] syscall=%d, os=%d", r1, os.Getpid())

	return nil
}

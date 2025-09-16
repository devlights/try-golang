//go:build unix

package main

import (
	"fmt"
	"log"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		// syscall(2)のための引数とポインタ情報
		trap    = uintptr(unix.SYS_GETCPU)
		cpu     uint32
		node    uint32
		ptrCpu  = uintptr(unsafe.Pointer(&cpu))
		ptrNode = uintptr(unsafe.Pointer(&node))

		// syscall(2)の結果
		r1    uintptr
		errno unix.Errno
		err   error
	)
	// getcpu(2)はブロッキングしない単純なシステムコールなのでRawSyscall()を使っても問題無い
	r1, _, errno = unix.RawSyscall(trap, ptrCpu, ptrNode, uintptr(0))
	if errno != unix.Errno(0) {
		err = errno
		return err
	}

	if int(r1) < 0 {
		return fmt.Errorf("getcpu syscall returned %d", r1)
	}

	log.Printf("CPU: %d, NUMA: %d", cpu, node)

	return nil
}

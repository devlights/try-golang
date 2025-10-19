//go:build unix && !arm64

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

var (
	gi int
)

func init() {
	log.SetFlags(0)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	os.Stdout.Sync()
	os.Stderr.Sync()

	var li int

	//
	// Fork
	//
	// unix.SYS_FORK は、CPUアーキテクチャが aarch64 (arm64) の場合は存在しないことに注意
	pid, _, errno := unix.RawSyscall(unix.SYS_FORK, 0, 0, 0)
	if errno != 0 {
		return fmt.Errorf("fork failed: %d", errno)
	}

	var err error
	switch pid {
	case 0:
		err = child(pid, &li)
	default:
		err = parent(pid)
	}

	if err != nil {
		return err
	}

	log.Printf("[%5d] li = %d, gi = %d", pid, li, gi)

	return nil
}

func child(pid uintptr, i *int) error {
	log.Printf("[%5d] This is the child process.", pid)

	for range 3 {
		log.Printf("[%5d] child processing...", pid)
		*i++
		gi++
		time.Sleep(time.Second)
	}

	return nil
}

func parent(pid uintptr) error {
	log.Printf("[%5d] This is the parent process.", pid)

	//
	// Join
	//
	log.Printf("[%5d] parent process WAIT started.", pid)
	defer log.Printf("[%5d] parent process WAIT done.", pid)

	var status unix.WaitStatus
	if _, err := unix.Wait4(int(pid), &status, 0, nil); err != nil {
		return fmt.Errorf("wait() failed: %w", err)
	}

	//
	// WIFEXITED(status) と等価
	//
	if status.Exited() {
		log.Printf("[%5d] Child process exited with status: %d", pid, status.ExitStatus())
	} else {
		log.Printf("[%5d] Child process did not exit normally: %d", pid, status.ExitStatus())
	}

	return nil
}

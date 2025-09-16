//go:build unix

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
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
	var (
		abs, _ = filepath.Abs(".")
		bin    = "/bin/ls"
		args   = []string{bin, "-lh", abs}
		envs   = []string{"PATH=/usr/bin:/bin"}
		attr   = syscall.ProcAttr{
			Dir:   abs,
			Env:   envs,
			Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		}

		pid    int
		err    error
		status syscall.WaitStatus
	)
	if pid, err = syscall.ForkExec(bin, args, &attr); err != nil {
		return fmt.Errorf("ForkExec() failed: %w", err)
	}

	log.Printf("[%5d] Child process started", pid)

	if _, err = syscall.Wait4(pid, &status, 0, nil); err != nil {
		return fmt.Errorf("Wait4() failed: %w", err)
	}

	if status.Exited() {
		log.Printf("[%5d] Child process exited with status: %d", pid, status.ExitStatus())
	} else {
		log.Printf("[%5d] Child process did not exit normally: %d", pid, status.ExitStatus())
	}

	return nil
}

//go:build !windows

// シグナルを送信する側です
//
// # 処理手順
//
//   - receiver プロセスを探して pid を取得
//   - 対象 pid に対して SIGTERM を送る
//
// REFERENCES:
//   - https://stackoverflow.com/questions/9030680/list-of-currently-running-process-in-go
package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

const (
	procName = "receiver"
)

var (
	appLog = log.New(os.Stderr, "[sender  ] >>> ", 0)
)

func main() {
	pid, err := find(procName)
	if err != nil {
		panic(err)
	}

	if pid == -1 {
		panic("receiver pid not found")
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	appLog.Printf("send SIGTERM to receiver(%d)", pid)
	err = proc.Signal(syscall.SIGTERM)
	if err != nil {
		panic(err)
	}
}

func find(name string) (int, error) {
	var (
		matches []string
		err     error
		pid     = -1
	)

	matches, err = filepath.Glob("/proc/*/exe")
	if err != nil {
		return pid, err
	}

	for _, f := range matches {
		real, err := os.Readlink(f)
		if err != nil && !errors.Is(err, os.ErrPermission) {
			return pid, err
		}

		if len(real) > 0 && strings.Contains(real, "receiver") {
			dir := filepath.Base(filepath.Dir(f))

			pid, err = strconv.Atoi(dir)
			if err != nil {
				return pid, err
			}

			appLog.Printf("receiver pid is %v\n", pid)
			break
		}
	}

	return pid, nil
}

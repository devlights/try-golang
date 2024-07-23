package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
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

	// forkする前にフラッシュしておく
	//
	// forkシステムコールが呼ばれる前にバッファに溜まっている場合
	// そのバッファが親プロセスと子プロセスの両方にコピーするため
	// 両方のプロセスが同じメッセージを出力する可能性があるため。
	//
	// 実際に以下をコメントアウトした状態で何回か実行すると
	// 結果が１つになったり、２つになったりする。
	os.Stdout.Sync()
	os.Stderr.Sync()

	//
	// forkシステムコールを呼び出し
	//
	// 通常の関数のように error が返ってくるのではなく
	// syscall.Errno が返ってくることに注意。
	pid, _, errno := unix.RawSyscall(unix.SYS_FORK, 0, 0, 0)
	if errno != 0 {
		return fmt.Errorf("fork failed: %d", errno)
	}

	switch pid {
	case 0:
		log.Printf("[%5d] This is the child process.", pid)
	default:
		log.Printf("[%5d] This is the parent process.", pid)
	}

	return nil
}

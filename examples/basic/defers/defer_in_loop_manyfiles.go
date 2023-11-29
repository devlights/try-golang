package defers

import (
	"fmt"
	"os"
	"runtime"

	"github.com/devlights/gomy/enumerable"
	"github.com/devlights/gomy/mem"
	"github.com/devlights/gomy/output"
)

// DeferInLoopManyFiles は、deferをループ内で利用したい場合のやり方についてのサンプルです。
// ループ内で大量のファイルを開いて defer で close しようとしている場合の対処について。
//
// REFERENCES::
//   - https://mattn.kaoriya.net/software/lang/go/20151212021608.htm
//   - https://stackoverflow.com/questions/45617758/defer-in-the-loop-what-will-be-better
func DeferInLoopManyFiles() error {
	// --------------------------------------------------------------
	// 基本的な動作については、 defer_in_loop.go にて記載している。
	// ここでは、ファイルハンドルをループ内でキャプチャさせたまま
	// defer で大量に登録した場合、メモリがどうなるかを検証する.
	// --------------------------------------------------------------
	dir, err := initDirectory()
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer os.RemoveAll(dir)
	output.Stdoutl("dir", dir)

	// ダメなパターン
	loopRange := enumerable.NewRange(1, 3000)
	memory := mem.NewMem(mem.Alloc(true), mem.TotalAlloc(false), mem.NumGC(true))
	err = badDefer(dir, loopRange, memory)
	if err != nil {
		return err
	}

	runtime.GC()

	// 良いパターン
	_, _ = loopRange.Reset()
	err = goodDefer(dir, loopRange, memory)
	if err != nil {
		return err
	}

	memory.Print("main end")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: defer_in_loop_manyfiles

	   [Name] "defer_in_loop_manyfiles"
	   dir                  /tmp/try-golang3234544150
	   badDefer - init      ----------------------------
	   Alloc                421 KiB
	   NumGC                0
	   [file count]         3001
	   badDefer - before gc ----------------------------
	   Alloc                2153 KiB
	   NumGC                0
	   badDefer - after gc  ----------------------------
	   Alloc                1140 KiB
	   NumGC                1
	   goodDefer - init     ----------------------------
	   Alloc                369 KiB
	   NumGC                2
	   [file count]         3001
	   goodDefer - before gc ----------------------------
	   Alloc                1819 KiB
	   NumGC                2
	   goodDefer - after gc ----------------------------
	   Alloc                369 KiB
	   NumGC                3
	   main end             ----------------------------
	   Alloc                369 KiB
	   NumGC                3


	   [Elapsed] 246.139489ms
	*/

}

func initDirectory() (string, error) {
	dir, err := os.MkdirTemp("", "try-golang")
	if err != nil {
		return "", err
	}

	return dir, nil
}

func badDefer(dir string, r enumerable.Range, memory mem.Mem) error {
	// 現在のメモリ量を出力しておく
	memory.Print("badDefer - init")

	// 大量にファイル作って defer に登録.
	// (関数スコープを作らない版)
	for r.Next() {
		file, err := os.CreateTemp(dir, fmt.Sprintf("try-golang-tmp-%02d", r.Current()))
		if err != nil {
			return err
		}

		// とりあえず適当にデータ入れておく
		err = os.WriteFile(file.Name(), []byte("helloworld"), os.ModePerm)
		if err != nil {
			return err
		}

		// defer 登録
		//   GoLand 使っている場合はループ内でのdeferの利用を検知してくれる
		//   サンプルなので抑止しておく
		//noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer file.Close()
	}

	output.Stdoutl("[file count]", r.Current()+1)

	// 現在のメモリ量を出力しておく
	memory.Print("badDefer - before gc")
	runtime.GC()
	memory.Print("badDefer - after gc")

	return nil
}

func goodDefer(dir string, r enumerable.Range, memory mem.Mem) error {
	// 現在のメモリ量を出力しておく
	memory.Print("goodDefer - init")

	// 大量にファイル作って defer に登録.
	// (関数スコープを作る版)
	for r.Next() {
		err := func() error {
			file, err := os.CreateTemp(dir, fmt.Sprintf("try-golang-tmp-%02d", r.Current()))
			if err != nil {
				return err
			}

			// とりあえず適当にデータ入れておく
			err = os.WriteFile(file.Name(), []byte("helloworld"), 0644)
			if err != nil {
				return err
			}

			// defer 登録
			//   GoLand 使っている場合はループ内でのdeferの利用を検知してくれる
			//   サンプルなので抑止しておく
			//noinspection GoUnhandledErrorResult,GoDeferInLoop
			defer file.Close()

			return nil
		}()

		if err != nil {
			return err
		}
	}

	output.Stdoutl("[file count]", r.Current()+1)

	// 現在のメモリ量を出力しておく
	memory.Print("goodDefer - before gc")
	runtime.GC()
	memory.Print("goodDefer - after gc")

	return nil
}

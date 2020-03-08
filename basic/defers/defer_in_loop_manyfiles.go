package defers

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/devlights/try-golang/lib/output"
	"github.com/devlights/try-golang/util/enumerable"
	"github.com/devlights/try-golang/util/mem"
)

// DeferInLoopManyFiles は、deferをループ内で利用したい場合のやり方についてのサンプルです。
// ループ内で大量のファイルを開いて defer で close しようとしている場合の対処について。
//
// REFERNCES::
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
}

func initDirectory() (string, error) {
	dir, err := ioutil.TempDir("", "try-golang")
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
		file, err := ioutil.TempFile(dir, fmt.Sprintf("try-golang-tmp-%02d", r.Current()))
		if err != nil {
			return err
		}

		// とりあえず適当にデータ入れておく
		err = ioutil.WriteFile(file.Name(), []byte("helloworld"), 0644)
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
			file, err := ioutil.TempFile(dir, fmt.Sprintf("try-golang-tmp-%02d", r.Current()))
			if err != nil {
				return err
			}

			// とりあえず適当にデータ入れておく
			err = ioutil.WriteFile(file.Name(), []byte("helloworld"), 0644)
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

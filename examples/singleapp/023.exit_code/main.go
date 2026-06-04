package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeout(rootCtx, 1*time.Second)
	)
	defer mainCxl()

	if err := run(mainCtx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	for i := range 10 {
		var (
			ctx, cxl = context.WithTimeout(ctx, 100*time.Millisecond)
			program  = os.Args[1]
			param    = strconv.Itoa(i)
			cmd      = exec.CommandContext(ctx, program, param)

			err      error
			exitCode int
			success  = true
		)
		defer cxl()

		if err = cmd.Run(); err != nil {
			//
			// 結果コードが 0、つまり、成功の場合は err は nil となる。
			//
			// 0以外の場合、errとして返ってくる。
			// この場合、err は *exec.ExitError となっているので
			// errors.As() で、変換できるか確認し、ExitCode() で取得する。
			//
			// Run() している最中に context がタイムアウトした場合
			// 結果コードは -1 となって返ってくる。
			//
			// Run() する前に context がタイムアウトした場合
			// context.DeadlineExceeded が返ってくる。
			//
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				exitCode = exitErr.ExitCode()
				success = exitErr.Success()
			} else {
				// 別のエラーの場合 (context.DeadlineExceeded など)
				return err
			}
		}

		fmt.Printf("実行引数: %d\t成功: %v\t結果コード: %d\n", i, success, exitCode)
	}

	return nil
}

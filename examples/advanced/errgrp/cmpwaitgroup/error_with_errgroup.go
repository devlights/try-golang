package cmpwaitgroup

import (
	"fmt"

	"github.com/devlights/try-golang/output"
	"github.com/devlights/try-golang/util/enumerable"
	"golang.org/x/sync/errgroup"
)

// ErrWithErrGroup は、拡張ライブラリ golang.org/x/sync/errgroup でエラー情報を呼び元に伝播させるサンプルです.
func ErrWithErrGroup() error {
	var (
		loopRange = enumerable.NewRange(1, 6)
		waitGrp   = errgroup.Group{}
	)

	// errgroup.Group は、sync.WaitGroup のように待ち合わせを行う機能に加えて
	// 発生したエラーを収集し、呼び元に返すことが可能となっている
	// 返してくれるエラーは、最初に発生したエラー情報となっている
	//
	// 利用方法は、sync.WaitGroup とは少し異なり Go(func() error) メソッドに
	// 非同期実行部分を渡して処理する形となっている. 内部で goroutine 化して呼び出してくれるので
	// 呼び元で go を付与する必要はない.
	//
	// 待ち合わせを実施したい箇所で、Wait() メソッドを呼び出すことにより非同期処理全部が完了するまで
	// 呼び元をブロックする。
	for loopRange.Next() {
		i := loopRange.Current()
		waitGrp.Go(func() error {
			prefix := fmt.Sprintf("[go func %02d]", i)
			output.Stderrl(prefix, "start")
			defer output.Stderrl(prefix, "end")

			err := randomErr(prefix)
			if err != nil {
				output.Stderrl(prefix, "\tERROR!!")
				return err
			}

			return nil
		})
	}

	// 複数の goroutine にて、複数のエラーが発生している場合でも取得できるのは最初に発生したエラーとなる
	if err := waitGrp.Wait(); err != nil {
		output.Stdoutl("[err]", err)
	}

	return nil
}

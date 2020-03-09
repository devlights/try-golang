package withcontext

import (
	"context"
	"fmt"

	"github.com/devlights/try-golang/output"
	"github.com/devlights/try-golang/util/enumerable"
	"golang.org/x/sync/errgroup"
)

// ErrGroupWithContext は、拡張ライブラリ golang.org/x/sync/errgroup で ctx.Context を含めた利用方法についてのサンプルです.
//
// https://pkg.go.dev/golang.org/x/sync/errgroup?tab=doc#example-Group-Parallel
func ErrGroupWithContext() error {
	// 利用するコンテキスト関連
	var (
		rootCtx           = context.Background()
		errGrp, errGrpCtx = errgroup.WithContext(rootCtx)
	)

	// その他の情報
	var (
		loopRange = enumerable.NewRange(1, 6)
	)

	// ----------------------------------------------------------------------------------------
	// errgroup.WithContext(ctx.Context) を利用することで、コンテキスト情報も管理することが可能となる
	// ここで取得した ctx.Context は、以下の場合にキャンセル状態となる。つまり、 <-ctx.Done() が通るようになる
	//   - どれかの非同期処理が最初に non-nil な戻り値を返したとき
	//   - 最初に Wait() が返ったとき
	// なので、非同期処理内でこのコンテキストを見張ることにより、どこかの処理でエラーが発生した場合に
	// まだ処理が始まっていない or 現在処理中の処理 をまとめてキャンセルすることができる
	// (現在処理中のものをキャンセルするためには、定周期で ctx.Done() を確認するポーリング処理を作り込む必要がある)
	// ----------------------------------------------------------------------------------------
	for loopRange.Next() {
		i := loopRange.Current()

		errGrp.Go(func() error {
			prefix := fmt.Sprintf("[go func %02d]", i)

			select {
			case <-errGrpCtx.Done():
				// だれかが初めにエラーを返した時点でこのコンテキストがキャンセルされる
				// main-goroutine側はWait() を呼び出しているため、この Wait() が return した
				// タイミングでもコンテキストはキャンセルされる.
				output.Stderrl(prefix, "CANCEL!!")
				return nil
			default:
				output.Stderrl(prefix, "start")
				defer output.Stderrl(prefix, "end")

				err := raiseErr(prefix)
				if err != nil {
					output.Stderrl(prefix, "\tERROR!!")
				}

				return err
			}
		})
	}

	if err := errGrp.Wait(); err != nil {
		output.Stdoutl("[err]", err)
	}

	return nil
}

func raiseErr(message string) error {
	return fmt.Errorf("raiseErr [%s]", message)
}

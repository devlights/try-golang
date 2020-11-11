package pipeline

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/devlights/gomy/output"
	"golang.org/x/sync/errgroup"
)

type (
	md5result struct {
		path     string
		checkSum [md5.Size]byte
		name     string
	}
)

// ErrGroupWithPipeline は、拡張ライブラリ golang.org/x/sync/errgroup でパイプライン処理を行っているサンプルです.
//
// https://pkg.go.dev/golang.org/x/sync/errgroup?tab=doc#example-Group-Pipeline
func ErrGroupWithPipeline() error {
	// 利用するコンテキスト関連
	var (
		rootCtx           = context.Background()
		errGrp, errGrpCtx = errgroup.WithContext(rootCtx)
	)

	var (
		filePathCh = make(chan string)
		md5Ch      = make(chan md5result)
	)

	// 1st ステージ
	// 配下の *.go ファイルをリストアップ
	errGrp.Go(func() error {
		defer close(filePathCh)
		return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(info.Name(), ".go") {
				filePathCh <- path
			}

			select {
			case <-errGrpCtx.Done():
				return errGrpCtx.Err()
			default:
				return nil
			}
		})
	})

	// 2nd ステージ
	// リストアップされたファイルを順次 md5 checksum していく
	// 10個のgoroutineを並行処理させる
	for i := 0; i < 10; i++ {
		goroutineIndex := i + 1
		errGrp.Go(func() error {
			var (
				name  = fmt.Sprintf("goroutine-%02d", goroutineIndex)
				count = 0
			)

			for p := range filePathCh {
				data, err := ioutil.ReadFile(p)
				if err != nil {
					return err
				}

				checksum := md5.Sum(data)
				result := md5result{
					path:     p,
					checkSum: checksum,
					name:     name,
				}

				select {
				case md5Ch <- result:
					count++
				case <-errGrpCtx.Done():
					return errGrpCtx.Err()
				}
			}

			return nil
		})
	}

	// 3rd ステージ
	// 1st, 2nd の処理完了を検知して結果用のチャネルである md5ch を閉じる
	go func() {
		_ = errGrp.Wait()
		close(md5Ch)
	}()

	// final ステージ
	// 結果出力
	for r := range md5Ch {
		cs := fmt.Sprintf("%x", r.checkSum)
		output.Stdoutl(r.name, cs, r.path)
	}

	// エラー判定
	// Wait() は、複数回呼んでも構わない.
	// 上の呼び出しは、処理の区切りを判定するために最初にエラーが返ったタイミング、もしくは、全部処理が終わったことを
	// 検知するためのもの。以下は、再度呼び出してエラーがあれば出力するためのもの
	if err := errGrp.Wait(); err != nil {
		output.Stdoutl("err", err)
	}

	return nil
}

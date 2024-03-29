package cmpwaitgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/devlights/gomy/enumerable"
	"github.com/devlights/gomy/output"
)

// ErrWithWaitGroup は、標準ライブラリ sync.WaitGroup でエラー情報を呼び元に伝播させるサンプルです.
func ErrWithWaitGroup() error {
	var (
		loopRange = enumerable.NewRange(1, 6)
		waitGrp   = sync.WaitGroup{}
		errorCh   = make(chan error)
		rnd       = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	// ----------------------------------------------------------------------------------------
	// sync.WaitGroup は、待ち合わせを担当するためのものであるため
	// 非同期処理側で発生したエラーを収集しておくような機能は持っていない
	// そのため、エラーが発生した場合の処理をユーザ側で作り込む必要がある
	// ----------------------------------------------------------------------------------------
	for loopRange.Next() {
		waitGrp.Add(1)

		go func(i int) {
			defer waitGrp.Done()

			prefix := fmt.Sprintf("[go func %02d]", i)
			output.Stderrl(prefix, "start")
			defer output.Stderrl(prefix, "end")

			err := randomErr(prefix, rnd)
			if err != nil {
				output.Stderrl(prefix, "\tERROR!!")
				errorCh <- err
			}

		}(loopRange.Current())
	}

	go func() {
		waitGrp.Wait()
		close(errorCh)
	}()

	for err := range errorCh {
		output.Stdoutl("[err]", err)
	}

	return nil
}

func randomErr(message string, rnd *rand.Rand) error {
	i := rnd.Intn(100)
	if i > 30 {
		return fmt.Errorf("randomErr [%d][%s]", i, message)
	}

	return nil
}

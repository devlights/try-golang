package times

import (
	"context"
	"math"
	"sync"
	"time"
)

// DoNDurations は、forループとtime.Before()を組み合わせて「Nの時間分何かを行う」という処理のサンプルです。
func DoNDurations() error {
	var (
		ctx = context.Background()
		wg  = sync.WaitGroup{}
		d   = 3 * time.Second
	)
	wg.Add(1)

	go func(ctx context.Context, d time.Duration) {
		defer wg.Done()

		// 指定された時間分、何かする
		var (
			end = time.Now().Add(d)
		)
		for time.Now().Before(end) {
			// 何か処理したことにする。
			for i := range 1000000 {
				_ = math.Sqrt(float64(i * i))
			}

			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(ctx, d)

	wg.Wait()

	return nil
}

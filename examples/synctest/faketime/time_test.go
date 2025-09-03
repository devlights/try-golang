package faketime

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/nalgeon/be"
)

const (
	SleepTime = 2 * time.Second
)

func TestNoFakeTime(t *testing.T) {
	start := time.Now()
	{
		// 普通に時間が流れるので指定時間スリープする
		time.Sleep(SleepTime)
	}
	elapsed := time.Since(start)

	be.Equal(t, elapsed, SleepTime) // 実際の結果は実行時に変化し、基本的に厳密な一致は出来ないことがほとんど
}

func TestUseFakeTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		{
			// Fake-Timeが使用されるためsynctest内では一瞬で完了する
			time.Sleep(SleepTime)
		}
		elapsed := time.Since(start)

		be.Equal(t, elapsed, SleepTime) // Fake-Timeにより結果が一致する
	})
}

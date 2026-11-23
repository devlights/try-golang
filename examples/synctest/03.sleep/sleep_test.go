package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/nalgeon/be"
)

// TestUseSyncTestSleep は、Go1.27で追加された synctest.Sleep のサンプルです。
// 一定時間後でタイムアウトする処理のテストをシミュレートしています。
//
// Go1.26までは、同じ挙動をするために
//
//	time.Sleep()
//	synctest.Wait()
//
// の2行を書く必要があった。これと全く同じ挙動が
//
//	synctest.Sleep()
//
// となっている。
func TestUseSyncTestSleep(t *testing.T) {
	const (
		TimeLimit = 10 * time.Second
	)
	var (
		fn = func(ctx context.Context, timedOut *bool) {
			select {
			case <-ctx.Done():
				return
			case <-time.After(TimeLimit):
				*timedOut = true
			}
		}
	)
	synctest.Test(t, func(t *testing.T) {
		var (
			ctx      = t.Context()
			timedOut = false
		)

		// 何かの非同期処理を実行していて所定時間でタイムアウトとする
		go fn(ctx, &timedOut)

		// 9秒経過時点にすすめる
		// (バブル内なので一瞬で時間が進む)
		synctest.Sleep(9 * time.Second)
		be.Equal(t, timedOut, false)

		// さらに999ミリ秒進める
		// Go1.26までの書き方で処理
		// (バブル内なので一瞬で時間が進む)
		{
			time.Sleep(999 * time.Millisecond)
			synctest.Wait()
		}
		be.Equal(t, timedOut, false)

		// 残り1ミリ秒進めてタイムアウトとなることを確認
		// (バブル内なので一瞬で時間が進む)
		synctest.Sleep(1 * time.Millisecond)
		be.Equal(t, timedOut, true)
	})
}

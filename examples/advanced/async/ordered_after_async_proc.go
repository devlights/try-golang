package async

import (
	"sort"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

// OrderedAfterAsyncProc -- chans.Enumerate() を使った非同期処理をした後に正しい順序に並び替えるサンプルです.
func OrderedAfterAsyncProc() error {
	type (
		result struct {
			index int
			value interface{}
		}
	)

	var (
		givenTime    = 1 * time.Second
		numGoroutine = 2
		items        = []interface{}{"hello", "world", "こんにちわ", "世界"}
		results      = make([]*result, 0, 0)
	)

	var (
		done  = make(chan struct{})
		outCh = make(chan interface{})
	)

	defer close(done)

	// 処理するのに t に指定された時間がかかる関数
	fn := func(item interface{}, t time.Duration) {
		<-time.After(t)
		output.Stdoutl("[処理]", item)
	}

	// パイプライン生成
	forEachCh := chans.ForEach(done, items...)
	enumerateCh := chans.Enumerate(done, forEachCh)
	doneChList := chans.FanOut(done, enumerateCh, numGoroutine, func(e interface{}) {
		if v, ok := e.(*chans.IterValue); ok {
			fn(v.Value, givenTime)
			outCh <- &result{
				index: v.Index,
				value: v.Value,
			}
		}
	})

	// 処理完了とともに出力用チャネルを閉じる
	go func() {
		defer close(outCh)
		<-chans.WhenAll(doneChList...)
	}()

	// 結果を吸い出し
	for v := range outCh {
		results = append(results, v.(*result))
	}

	// 正しい順序に並び替え
	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	// 最終結果を出力
	output.StdoutHr()
	for _, v := range results {
		output.Stdoutl("[最終結果]", v.value)
	}

	return nil
}

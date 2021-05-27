package async

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/devlights/gomy/chans"
)

// TakeFirst10Items -- 最初の１０個のみを取得するサンプルです
func TakeFirst10Items() error {
	var (
		rootCtx         = context.Background()
		mainCtx, cancel = context.WithCancel(rootCtx)
	)

	defer cancel()

	// 乱数を返す関数
	randomInt := func() interface{} {
		return rand.Int()
	}

	// 乱数を延々と返すチャネル生成
	repeatCh := chans.RepeatFn(mainCtx.Done(), randomInt)

	// 最初の１０件のみ取得するチャネル生成
	takeFirst10ItemCh := chans.Take(mainCtx.Done(), repeatCh, 10)

	// 出力
	for v := range takeFirst10ItemCh {
		fmt.Println(v)
	}

	return nil
}

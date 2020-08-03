package chapter01

import (
	"sync"

	"github.com/devlights/gomy/output"
)

// RaceConditionFixWithMutex -- sync.Mutex を使って ロック をかけて値の更新をすることで競合状態を回避するサンプルです.
func RaceConditionFixWithMutex() error {
	var (
		memoryAccess sync.Mutex
		data         int
	)

	go func() {
		defer memoryAccess.Unlock()
		memoryAccess.Lock()

		data++
	}()

	// mutex を使って、データの更新時にロックがかかるので
	// if 評価時と結果出力時にメモリ上の値が異なっているという状態は回避できるようになった。
	// しかし、この場合でも 値 自体は非決定的であることは変わらないので
	// 以下の if 分岐は、どちらを通るかは、動かす度に変化する可能性がある
	defer memoryAccess.Unlock()
	memoryAccess.Lock()

	if data == 0 {
		output.Stdoutl("[result]", "data is zero")
	} else {
		output.Stdoutl("[reslt]", "data is non-zero")
	}

	return nil
}

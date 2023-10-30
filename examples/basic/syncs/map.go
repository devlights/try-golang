package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/syncs/maps"
)

// Map は、sync.Mapの使い方とそれをラップした型の紹介です。
//
// # REFERENCES
//   - https://hjr265.me/blog/synchronization-constructs-in-go-standard-library/
//   - https://pkg.go.dev/sync@go1.21.3#Map
func Map() error {
	//
	// Goでは、標準のマップはスレッドセーフではない.
	// スレッドセーフなマップは sync.Map として提供されている.
	// が、この実装は map[any]any となっている.
	//

	var (
		m1  sync.Map
		wg1 sync.WaitGroup
	)

	wg1.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg1.Done()
			m1.Store(i, true)
		}()
	}

	wg1.Wait()
	m1.Range(func(key any, value any) bool {
		var (
			k = key.(int)
			v = value.(bool)
		)

		output.Stdoutf("[m1]", "%d:%v\n", k, v)
		return true
	})

	if v, ok := m1.Load(8); ok {
		if vv, okok := v.(bool); okok {
			output.Stdoutl("[m1.Load(8)]", vv)
		}
	}

	output.StdoutHr()

	//
	// Go 1.18 よりジェネリクスが利用できるようになっているので
	// sync.Mapをラップしたジェネリックな型を用意することが楽になった。
	//

	var (
		m2  maps.SafeMap[int, bool]
		wg2 sync.WaitGroup
	)

	wg2.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg2.Done()
			m2.Store(i, true)
		}()
	}

	wg2.Wait()
	m2.Range(func(key int, value bool) bool {
		output.Stdoutf("[m2]", "%d:%v\n", key, value)
		return true
	})

	if v, ok := m2.Load(8); ok {
		output.Stdoutl("[m2.Load(8)]", v)
	}

	return nil
}

package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/syncs/maps"
)

// UseMap は、sync.Mapの使い方とそれをラップした型の紹介です。
//
// # REFERENCES
//   - https://hjr265.me/blog/synchronization-constructs-in-go-standard-library/
//   - https://pkg.go.dev/sync@go1.21.3#Map
func UseMap() error {
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_map

	   [Name] "syncs_use_map"
	   [m1]                 5:true
	   [m1]                 4:true
	   [m1]                 9:true
	   [m1]                 7:true
	   [m1]                 8:true
	   [m1]                 1:true
	   [m1]                 0:true
	   [m1]                 2:true
	   [m1]                 3:true
	   [m1]                 6:true
	   [m1.Load(8)]         true
	   --------------------------------------------------
	   [m2]                 7:true
	   [m2]                 5:true
	   [m2]                 6:true
	   [m2]                 8:true
	   [m2]                 9:true
	   [m2]                 2:true
	   [m2]                 4:true
	   [m2]                 0:true
	   [m2]                 3:true
	   [m2]                 1:true
	   [m2.Load(8)]         true


	   [Elapsed] 507.49µs
	*/

}

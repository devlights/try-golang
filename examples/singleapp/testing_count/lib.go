package lib

import "time"

func Fn(count int) int64 {
	var (
		v int64
	)

	for i := 0; i < count; i++ {
		go func() {
			v += 1
		}()
	}

	time.Sleep(10 * time.Millisecond)

	return v
}

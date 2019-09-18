package time_

import (
	"fmt"
	"time"
)

// TimeSince は、 time.Since() のサンプルです.
func TimeSince() error {
	// ------------------------------------------------------------
	// time.Since
	//
	// 現在の時刻と指定された時刻の差を返してくれる.
	//
	// start := time.Now()
	// elappsed := time.Since(start)
	//
	// は
	//
	// start := time.Now()
	// elappsed := time.Now().Sub(start)
	//
	// とほぼ同じ.
	// ------------------------------------------------------------
	var (
		start    = time.Now()
		elappsed time.Duration
	)

	time.Sleep(3 * time.Second)

	elappsed = time.Since(start)
	fmt.Printf("Elappsed: %v\n", elappsed)

	return nil
}

package times

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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_since

	   [Name] "time_since"
	   Elappsed: 3.001573634s


	   [Elapsed] 3.001628754s
	*/

}

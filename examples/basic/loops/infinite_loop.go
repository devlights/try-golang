package loops

import (
	"context"
	"fmt"
	"time"
)

// InfiniteLoop は、無限ループのサンプルです.
func InfiniteLoop() error {
	var (
		mainCtx, mainCxl = context.WithCancel(context.Background())
		procCtx, procCxl = context.WithTimeout(mainCtx, 3*time.Second)
	)

	defer mainCxl()
	defer procCxl()

	// Go では 無限ループ は以下のように for {} とだけ書く
LOOP:
	for {
		select {
		case <-time.After(200 * time.Millisecond):
			fmt.Print(".")
		case <-procCtx.Done():
			break LOOP
		}
	}

	fmt.Println("")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_infinite_loop

	   [Name] "loops_infinite_loop"
	   ..............


	   [Elapsed] 3.000575495s
	*/

}

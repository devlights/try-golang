// signal.Ignore() のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/os/signal@go1.19.3#Ignore
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// SIGINTを無視するよう設定
	signal.Ignore(os.Interrupt)

	fmt.Println("SIGINTを無視に設定しています...5秒経つとアプリケーションは終了します.")
	<-time.After(5 * time.Second)
}

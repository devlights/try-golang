package loops

import "github.com/devlights/gomy/output"

// ChannelLoop は、チャネルをループさせる場合のサンプルです.
func ChannelLoop() error {
	var (
		ch    = make(chan string)
		quit  = make(chan struct{})
		items = []string{
			"go",
			"java",
			"dotnet",
			"python",
			"flutter",
		}
	)

	go func(ch chan<- string) {
		defer close(ch)

		for _, v := range items {
			ch <- v
		}
	}(ch)

	go func(quit chan<- struct{}, ch <-chan string) {
		defer close(quit)

		// チャネルを foreach ループする場合, インデックスは付かない
		for v := range ch {
			output.Stdoutl("", v)
		}
	}(quit, ch)

	<-quit

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_channel_loop

	   [Name] "loops_channel_loop"
	   go
	   java
	   dotnet
	   python
	   flutter


	   [Elapsed] 248.37µs
	*/

}

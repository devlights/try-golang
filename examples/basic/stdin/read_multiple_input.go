package stdin

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// ReadMultipleInput -- fmt.Scan() を利用して標準入力から複数の入力値を受け取るサンプルです。
//
// REFERENCES:
//   - https://dev.to/itnext/go-from-the-beginning-reading-user-input-i79
//   - https://stackoverflow.com/questions/15413469/how-to-make-fmt-scanln-read-into-a-slice-of-integers
func ReadMultipleInput() error {
	var (
		messages [3]string
		err      error
	)

	// fmt.Scan() は渡した引数を変更するので ポインタ で渡す必要がある
	// fmt.Scan() は空白を見つけると探索を中止する
	// 引数は可変長となっているので、以下のように指定することが出来る
	fmt.Print("ENTER 3 WORDS: ")
	_, err = fmt.Scan(&messages[0], &messages[1], &messages[2])
	if err != nil {
		return err
	}

	output.Stdoutf("[fmt.Scan][01]", "messages=%v\n", messages)

	// ループして一つずつ取得してもよい
	// この場合、よく利用する for i, v := range messages {} で処理すると v の値は
	// Go の ループ では参照を使いまわししているのでうまくいかない。
	// 以下のように for i := range messages {} で添字を使って処理する。
	fmt.Print("ENTER 3 WORDS: ")

	messages = [3]string{}
	for i := range messages {
		_, err = fmt.Scan(&messages[i])
		if err != nil {
			return err
		}
	}

	output.Stdoutf("[fmt.Scan][02]", "messages=%v\n", messages)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: stdin_read_multiple_input

	   [Name] "stdin_read_multiple_input"
	   ENTER 3 WORDS: java dotnet golang
	   [fmt.Scan][01]       messages=[java dotnet golang]
	   ENTER 3 WORDS: java dotnet golang
	   [fmt.Scan][02]       messages=[java dotnet golang]


	   [Elapsed] 8.880712932s
	*/

}

package logging

import (
	"log"
)

// Prefix は、log.Prefix() と log.SetPrefix() の挙動を確認するサンプルです。
func Prefix() error {
	// ----------------------------------------------------------------
	// log.Prefix()とlog.SetPrefix()について
	//
	// log.Prefix() は、現在設定されているログの接頭辞を取得することが出来る.
	// log.SetPrefix()は、新たにログの接頭辞を設定することが出来る.
	//
	// デフォルトでは、log.Prefix() には 空文字 が設定されている.
	//
	// REFERENCES::
	//   https://golang.org/log/#SetPrefix
	//   https://golang.org/log/#Prefix
	// ----------------------------------------------------------------
	// 一時的にログの出力フォーマットを変更
	origFlags := log.Flags()
	log.SetFlags(0)
	defer log.SetFlags(origFlags)

	// デフォルトの Prefix は 空
	origPrefix := log.Prefix()
	log.Printf("デフォルトのPrefix [%s]", origPrefix)
	defer log.SetPrefix(origPrefix)

	// 新たに接頭辞を設定
	log.SetPrefix("[mylog] ")

	// 接頭辞がついていることを確認
	log.Printf("接頭辞付き")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: logging_prefix

	   [Name] "logging_prefix"
	   デフォルトのPrefix []
	   [mylog] 接頭辞付き


	   [Elapsed] 25.6µs
	*/

}

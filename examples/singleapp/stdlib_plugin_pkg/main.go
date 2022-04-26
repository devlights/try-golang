package main

import (
	"log"
	"plugin"
)

func exitOnErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var (
		plg *plugin.Plugin
		sym plugin.Symbol
		err error
	)

	// プラグインをオープンして
	plg, err = plugin.Open("lib/lib.so")
	exitOnErr(err)

	// シンボルを取得
	sym, err = plg.Lookup("Fn")
	exitOnErr(err)

	// シンボルは 対象となる値 (変数とか関数) へのポインタとなっている
	// 関数の場合は sym.(func()) で利用できるが、変数の場合は *(sym.(*int)) のようにする必要がある点に注意
	if fn, ok := sym.(func(string)); ok {
		fn("hello world")
	}
}

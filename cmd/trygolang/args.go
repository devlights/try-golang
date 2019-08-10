package main

import "flag"

// Args は、プログラム引数の値を持つ構造体です１
type Args struct {
	OneTime   bool
	ShowNames bool
}

// NewArgs は、Argsのコンストラクタ関数です
func NewArgs() *Args {
	return new(Args)
}

// Parse は、コマンドライン引数を解析しパッケージ変数に格納します
func (a *Args) Parse() {
	flag.BoolVar(&a.OneTime, "onetime", false, "run only one time")
	flag.BoolVar(&a.ShowNames, "list", false, "show all example names")

	flag.Parse()
}

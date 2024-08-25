package main

import "flag"

// Args は、プログラム引数の値を持つ構造体です
type Args struct {
	// 一度だけ実行するかどうか
	OneTime bool
	// 実行可能な名前を表示するかどうか
	ShowNames bool
	// 実行するサンプル名
	ExampleName string
	// VSCodeで実行されているかどうか
	RunWithVsCode bool
}

// NewArgs は、Argsのコンストラクタ関数です
func NewArgs() *Args {
	return new(Args)
}

// Parse は、コマンドライン引数を解析しパッケージ変数に格納します
func (a *Args) Parse() {
	flag.BoolVar(&a.OneTime, "onetime", false, "run only one time")
	flag.BoolVar(&a.ShowNames, "list", false, "show all example names")
	flag.StringVar(&a.ExampleName, "example", "", "name of example")
	flag.BoolVar(&a.RunWithVsCode, "vscode", false, "run with vscode")

	flag.Parse()
}

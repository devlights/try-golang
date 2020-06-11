package main

type (
	// Command -- 何かを実行するコマンドを表します.
	Command interface {
		// Run -- 実行します.
		Run() error
	}
)

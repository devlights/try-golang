package command

type (
	// Cmd -- 何かを実行するコマンドを表します.
	Cmd interface {
		// Run -- 実行します.
		Run() error
	}
)

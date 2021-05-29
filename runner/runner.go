package runner

type (
	// Runner -- 何かを実行するコマンドを表します.
	Runner interface {
		// Run -- 実行します.
		Run() error
	}
)

package main

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		v any = 100
	)

	// そのままにしていると v は利用されていないことになるため
	// コンパイルエラー (v declared and not usedcompilerUnusedVar) となる。
	// 一時的に抑止したい場合は _ に代入しておく
	_ = v

	return nil
}

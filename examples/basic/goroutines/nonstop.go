package goroutines

import "github.com/devlights/gomy/output"

// NonStop -- ゴルーチンを待ち合わせ無しで走らせるサンプルです.
//
// 投げっぱなしのゴルーチンを作る場合に使います。
// 通常待ち合わせ無しの非同期処理は行うべきではありません。
func NonStop() error {
	go func() {
		output.Stdoutl("[goroutine] ", "This line may not be printed")
	}()

	// 上記のゴルーチンは待ち合わせをしていないので出力されない可能性がある。
	// （出力する前にメインゴルーチンが終わる可能性がある)

	return nil
}

package errs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/devlights/gomy/output"
)

// IsAndAs -- errors.Is(), errors.As() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/errors@go1.19
//   - https://zenn.dev/msksgm/articles/20220325-unwrap-errors-is-as
func IsAndAs() error {
	_, err := os.Open("notexists")
	wrapped := fmt.Errorf("wrap err %w", err)

	// errors.Is() は、同じ値かどうかを判定してくれて、かつ、自動でUnwrapもしてくれる
	// 予め、エラーの値が存在している場合は errors.Is() で確認
	if errors.Is(wrapped, err) {
		output.Stdoutl("[errors.Is]", "wrapped を Unwrap していくと err がある")
	}

	// errors.Is() は、Unwrapしてくれることを除くと以下と同じ意味となる
	if wrapped == err {
		output.Stdoutl("[errors.Is]", "Unwrapしていないのでここは入らない")
	}

	// 手動で errors.Unwrap() を呼び出すと同じ動作となる
	if errors.Unwrap(wrapped) == err {
		output.Stdoutl("[errors.Is]", "errors.Unwrapを呼び出して手動でUnwrapした場合")
	}

	// errors.As() は、代入可能かどうかを判定してくれて、かつ、自動でUnwrapもしてくれる。
	// 実行時に発生したエラーが特定の型のエラーであるかどうかを確認したい場合は errors.As() で確認。
	// ポインタを渡す必要があることに注意。
	//
	// Go 1.19 にて go vet に警告が追加されたので、以下を
	//   var pErr fs.PathError
	// としてしまっている場合は、警告が表示されるようになった。(errorsas)
	var pErr *fs.PathError
	if errors.As(wrapped, &pErr) {
		output.Stdoutl("[errors.As]", "wrapped を Unwrap していくと pErr に代入可能")
	}

	// errors.As() は、Unwrapしてくれることを除くと以下と同じ意味となる
	if _, ok := wrapped.(*fs.PathError); ok {
		output.Stdoutl("[errors.As]", "Unwrapしていないのでここは入らない")
	}

	// 手動で errors.Unwrap() を呼び出すと同じ動作となる
	if _, ok := errors.Unwrap(wrapped).(*fs.PathError); ok {
		output.Stdoutl("[errors.As]", "errors.Unwrapを呼び出して手動でUnwrapした場合")
	}

	return nil
}

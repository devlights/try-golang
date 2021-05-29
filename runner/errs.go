package runner

import (
	"fmt"

	"github.com/devlights/try-golang/mapping"
)

type (
	// ExecError -- 実行時エラーを表します.
	ExecError struct {
		Name mapping.ExampleKey // 名称
		Err  error              // エラー
	}
)

func (e *ExecError) Error() string {
	return fmt.Sprintf("%s (%s)", e.Err.Error(), e.Name)
}

// Unwrap -- 内部エラーを返します.
func (e *ExecError) Unwrap() error {
	return e.Err
}

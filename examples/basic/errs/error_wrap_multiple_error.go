package errs

import (
	"errors"
	"fmt"

	"github.com/devlights/gomy/output"
)

type _wmErr struct {
	m string
}

type _wmErr2 struct {
}

func (me *_wmErr) Error() string {
	return me.m
}

func (me *_wmErr2) Error() string {
	return ""
}

// WrapMultipleError は、Go 1.20 で導入された %w を複数指定できるようになった機能のサンプルです。
//
// # REFERENCES
//   - https://future-architect.github.io/articles/20230126a/
//   - https://tip.golang.org/doc/go1.20#errors
func WrapMultipleError() error {
	//
	// Go 1.20 で、 %w を複数指定できるようになった
	//
	var (
		e1 = errors.New("error 1")
		e2 = errors.New("error 2")
		e3 = &_wmErr{m: "error 3"}
	)

	e4 := fmt.Errorf("%w,%w,%w", e1, e2, e3)
	output.Stdoutf("[e4]", "%T: %v\n", e4, e4)

	//
	// errors.Is でちゃんと判定される
	//
	output.Stdoutl("[Is(e4, e2)]", errors.Is(e4, e2))

	//
	// errors.As でちゃんと判定される
	//
	var wme *_wmErr
	var wme2 *_wmErr2
	output.Stdoutl("[As(e4, &wme) ]", errors.As(e4, &wme))
	output.Stdoutl("[As(e4, &wme2)]", errors.As(e4, &wme2))

	return nil
}

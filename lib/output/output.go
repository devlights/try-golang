// output は、サンプル実行時の出力部分を抽象化するパッケージです.
package output

import (
	"fmt"
	"io"
	"os"
)

// デフォルトの 標準出力と標準エラー出力 に利用する io.Writer
var (
	defaultWriter    io.Writer = os.Stdout
	defaultErrWriter io.Writer = os.Stderr
)

// Writer は、現在設定されている標準出力用のio.Writerを返します.
func Writer() io.Writer {
	return defaultWriter
}

// ErrWriter は、現在設定されている標準エラー出力用のio.Writerを返します.
func ErrWriter() io.Writer {
	return defaultErrWriter
}

// SetWriter は、標準出力用のio.Writerを設定します.
func SetWriter(w io.Writer) {
	defaultWriter = w
}

// SetErrWriter は、標準エラー出力用のio.Writerを設定します.
func SetErrWriter(w io.Writer) {
	defaultErrWriter = w
}

// Stdoutl は、指定された接頭辞と値を標準出力に出力します.
func Stdoutl(prefix string, values ...interface{}) {
	_pl(Writer(), prefix, values...)
}

// Stderrl は、指定された接頭辞と値を標準エラー出力に出力します.
func Stderrl(prefix string, values ...interface{}) {
	_pl(ErrWriter(), prefix, values...)
}

// Stdoutf は、指定された接頭辞と書式付きの値を標準出力に出力します.
func Stdoutf(prefix string, format string, values ...interface{}) {
	_pf(Writer(), prefix, format, values...)
}

// Stderrf は、指定された接頭辞と書式付きの値を標準出力に出力します.
func Stderrf(prefix string, format string, values ...interface{}) {
	_pf(ErrWriter(), prefix, format, values...)
}

func _pf(w io.Writer, prefix string, format string, values ...interface{}) {
	if prefix != "" {
		s := fmt.Sprintf(format, values...)
		_, _ = fmt.Fprintf(w, "%s %s", prefix, s)
	} else {
		_, _ = fmt.Fprintf(w, format, values...)
	}
}

func _pl(w io.Writer, prefix string, values ...interface{}) {
	if prefix != "" {
		s := fmt.Sprintln(values...)
		_, _ = fmt.Fprint(w, prefix, " ", s)
	} else {
		_, _ = fmt.Fprintln(w, values...)
	}
}

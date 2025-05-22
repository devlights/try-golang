package flags

import (
	"flag"
	"fmt"

	"github.com/devlights/gomy/output"
)

type (
	varsFlag []string
)

func (me *varsFlag) String() string {
	return fmt.Sprint(*me)
}

func (me *varsFlag) Set(v string) error {
	*me = append(*me, v)
	return nil
}

var (
	// varsFlag は flag.Value を実装している
	_ flag.Value = (*varsFlag)(nil)
)

// Var2 は、flagパッケージのflag.Var()のサンプルです。
//
// コマンドライン引数にて同じオプションを複数回指定された場合に対応できる
// カスタムフラグを定義して、値を読み取っています。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.24.3
func Var2() error {
	var (
		fs   = flag.NewFlagSet("vars", flag.ExitOnError)
		vars varsFlag
	)
	fs.Var(&vars, "v", "文字列値。複数指定可能。")

	var (
		opts = []string{
			"-v", "hello",
			"-v", "world",
			"-v", "へろー",
			"-v", "ワールド",
		}
		err error
	)
	err = fs.Parse(opts)
	if err != nil {
		return err
	}

	var (
		ch = make(chan string)
	)
	go func(ch chan<- string, vars []string) {
		defer close(ch)
		for _, item := range vars {
			ch <- item
		}
	}(ch, vars)

	for item := range ch {
		output.Stdoutl("[vars]", item)
	}

	return nil
}

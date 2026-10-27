package main

import (
	"context"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"os"
)

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(_ context.Context) error {
	var (
		buf []byte
		err error
	)
	buf, err = io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	//
	// 標準ライブラリの go/types パッケージにある types.Eval を使うと、文字列の式を1回の呼び出しで評価出来る。
	// 変数付きで数式評価とかしたい場合は [gval](https://github.com/PaesslerAG/gval) を使ったほうが良い。
	// 定数式だけなら、types.Eval で十分。
	//
	var (
		fset *token.FileSet = token.NewFileSet() // ファイルセットは空
		pkg  *types.Package = nil                // nilを指定するとUniverseスコープとなる。今回は定数のみなのでnilで良い。
		pos  token.Pos      = token.NoPos        // token.Pos のゼロ値は token.NoPos だが、サンプルなので意図的に指定
		expr                = string(buf)
		tv   types.TypeAndValue
	)
	tv, err = types.Eval(fset, pkg, pos, expr)
	if err != nil {
		return err
	}

	fmt.Println(tv.Value.ExactString())

	return nil
}

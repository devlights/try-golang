package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// Numbers は、数値のフォーマットについてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/fmt
func Numbers() error {
	// '+' を付与すると常に符号が表示される
	{
		x := 123
		fmt.Printf("%+d\n", x)
		fmt.Printf("%+d\n", x*-1)
	}
	output.StdoutHr()

	// 指定した値分だけ右側をスペースで埋める
	{
		x := 123
		fmt.Printf("'%-10d'\n", x)
		fmt.Printf("'%-10d'\n", x*-1)
	}
	output.StdoutHr()

	// 別の進数で出力
	{
		x := 0xff
		fmt.Printf("%d\n", x)
		fmt.Printf("%#b\n", x)
		fmt.Printf("%#x\n", x)
		fmt.Printf("%#o\n", x)
	}
	output.StdoutHr()

	// 要素毎にスペースを開けて出力
	{
		fmt.Printf("% x\n", [...]byte{253, 254, 255})
		fmt.Printf("% x\n", []byte{250, 251, 252})
		fmt.Printf("% x\n", "hello")
	}
	output.StdoutHr()

	// 0埋めして出力
	{
		fmt.Printf("'%08d'\n", 0xf0)
		fmt.Printf("'%08x'\n", 0xff)
		fmt.Printf("'%08b'\n", 0x0f)
	}

	return nil
}

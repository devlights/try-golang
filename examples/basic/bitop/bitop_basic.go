package bitop

import (
	"github.com/devlights/gomy/output"
)

// Basic -- 基本的なビット操作のサンプルです.
//
// REFERENCES:
//   - 書籍「プログラミング言語 Go」--  P.60
//   - https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
func Basic() error {
	// -----------------------------------------------------
	// ビット演算は集合演算と思うと理解しやすい.
	//
	// &  -- AND    , 共通部分
	// |  -- OR     , 和
	// ^  -- XOR    , 対称差 (どちらか一方しか持っていないもの)
	// &^ -- AND NOT, 差 (自分しか持っていないもの)、ビットクリア
	// -----------------------------------------------------

	var (
		x uint8 = 1<<0 | 1<<1 | 1<<5 // 00100011
		y uint8 = 1<<1 | 1<<3        // 00001010
	)

	// フォーマットのverbに %b を指定すると2進数表記となる
	// さらに 08 を付与して、8桁表示でゼロ埋めするように指示
	//
	// 補足) %#08b とすると, 先頭に 0b を付与して出力してくれる
	output.Stdoutf("[x]", "%08b\n", x)
	output.Stdoutf("[y]", "%08b\n", y)

	output.StdoutHr()

	// & は共通部分を取り出す。両方 1 のものが残る
	output.Stdoutf("[x&y ][共通部分]", "%08b\n", x&y)

	// | は和を取り出す。どちらか 1 のものが残る
	output.Stdoutf("[x|y ][和　　　]", "%08b\n", x|y)

	// ^ は対称差を取り出す。どちらか片方しか 1 のものが残る.
	output.Stdoutf("[x^y ][対称差　]", "%08b\n", x^y)

	// &^ は差を取り出す。 &^ はビットクリアともいう.
	// x&^yは、xの個々のビットは対応するyのビットが1の場合に0となる。
	// それ以外は、対応するxのビットがそのまま残る.
	//
	// 例) logパッケージの、デフォルトのフラグ設定をクリアする際に
	//        log.SetFlags(log.Flags &^ log.LstdFlags)
	//     としてビットクリアすると、一発でクリアできる
	//     (デフォルトで設定されているフラグ LstdFlags なので)
	output.Stdoutf("[x&^y][差　　　]", "%08b\n", x&^y)

	output.StdoutHr()

	// 特定のビットが立っているか検査
	if x&(1<<5) != 0 {
		output.Stdoutl("[x&(1<<5)]", "ON")
	}

	output.StdoutHr()

	// 特定のビットを落とす
	x &^= (1 << 5)
	output.Stdoutf("[x &^= (1 << 5)]", "%08b\n", x)

	// 特定のビットを立てる
	x |= (1 << 5)
	output.Stdoutf("[x != (1 << 5)]", "%08b\n", x)

	output.StdoutHr()

	// 特定のビットをトグル
	x ^= (1 << 7)
	output.Stdoutf("[x ^= (1 << 7)]", "%08b\n", x)
	x ^= (1 << 7)
	output.Stdoutf("[x ^= (1 << 7)]", "%08b\n", x)

	output.StdoutHr()

	// 自分の全ビットを反転
	output.Stdoutf("[^x]", "%08b\n", ^x)

	return nil
}

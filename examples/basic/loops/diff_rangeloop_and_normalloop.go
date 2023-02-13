package loops

import "github.com/devlights/gomy/output"

// DiffRangeLoopAndNormalLoop は、 range ループと通常のループの違いについてのサンプルです.
func DiffRangeLoopAndNormalLoop() error {
	var (
		s1 = []int{1, 2, 3}
	)

	//
	// range ループは、ループ時に元の値のコピーを作成してそれをループする。
	// 以下の処理は、ループの度に元のスライスに要素を追加しているが
	// rangeループでコピーしたスライスは変化が無いため、ループは３回で終了し
	// 元のスライスの要素数も増えて６となる。
	//
	for range s1 { // ループが始まる直前に内部的にコピーが取られ、そのコピーを元にループするイメージ
		s1 = append(s1, 99)
	}
	output.Stdoutl("[range-loop]", s1)

	//
	// 逆に、ループの度に継続条件が変化するため
	// 以下のループは永遠に終わらない。
	//
	/*
		var (
			s2 = []int{1, 2, 3}
		)
		for i := 0; i < len(s2); i++ {
			s2 = append(s2, 99)
		}
		output.Stdoutl("[ここには到達しない]", s2)
	*/

	return nil
}

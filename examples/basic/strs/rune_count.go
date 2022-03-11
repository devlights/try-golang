package strs

import (
	"unicode/utf8"

	"github.com/devlights/gomy/output"
)

// RuneCount -- utf8.RuneCountInString() のサンプルです。
//
// # REFERECES
//   - https://qiita.com/wifecooky/items/c1a06e8639a0a6c6f11f
//   - https://qiita.com/tana6/items/72f3064d1fb1f65f4470
func RuneCount() error {
	var (
		s1 = "h"
		s2 = "あ"
		s3 = "😺"
		s4 = "🧑‍🤝‍🧑"   //lint:ignore ST1018 ok
		s5 = "👨‍👩‍👧‍👦" //lint:ignore ST1018 ok
	)

	var (
		c1 = utf8.RuneCountInString(s1)
		c2 = utf8.RuneCountInString(s2)
		c3 = utf8.RuneCountInString(s3)
		c4 = utf8.RuneCountInString(s4)
		c5 = utf8.RuneCountInString(s5)
	)

	output.Stdoutf("[1]", "%q\t%d rune(s)\t%d byte(s)\n", s1, c1, len(s1))
	output.Stdoutf("[2]", "%q\t%d rune(s)\t%d byte(s)\n", s2, c2, len(s2))
	output.Stdoutf("[3]", "%q\t%d rune(s)\t%d byte(s)\n", s3, c3, len(s3))
	output.Stdoutf("[4]", "%q\t%d rune(s)\t%d byte(s)\n", s4, c4, len(s4))
	output.Stdoutf("[5]", "%q\t%d rune(s)\t%d byte(s)\n", s5, c5, len(s5))

	return nil
}

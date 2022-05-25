package exp_maps

import (
	"github.com/devlights/gomy/output"
	"golang.org/x/exp/maps"
)

// ExpMaps -- Go 1.18 リリース時には含まれなかったジェネリクス対応 汎用map処理が定義されている golang.org/x/exp/maps パッケージのサンプルです。
func ExpMaps() error {
	var (
		m1 = map[string]int{"hello": 100, "world": 101}
		m2 = map[string]int{"world": 101, "hello": 100}
		k  []string
		v  []int
	)

	k = maps.Keys(m1)
	v = maps.Values(m1)

	output.Stdoutl("[maps.Keys]", k, v)
	output.Stdoutl("[maps.Values]", maps.Equal(m1, m2))

	maps.Clear(m1)
	output.Stdoutl("[maps.Clear]", m1, m2)

	return nil
}

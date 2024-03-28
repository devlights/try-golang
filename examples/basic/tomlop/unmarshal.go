package tomlop

import (
	"github.com/devlights/gomy/output"
	"github.com/pelletier/go-toml/v2"
)

// Unmarshal は、TOMLファイルの内容を読み込むサンプルです.
//
// # REFERENCES
//   - https://www.meetgor.com/golang-config-file-read/
//   - https://github.com/pelletier/go-toml
//   - https://ja.wikipedia.org/wiki/TOML
func Unmarshal() error {
	const (
		src = `[values]
value1 = 999
value2 = 'hello world'
value3 = false
value4 = ['one', 'two', 'three']

[author]
name = 'devlights'

[[persons]]
name = 'one'
age  = 30

[[persons]]
name = 'two'
# age は 省略

[[persons]]
name = 'three'
age  = 99`
	)

	type (
		ValuesSection struct {
			Value1 int      `toml:"value1"` // 数値
			Value2 string   `toml:"value2"` // 文字列
			Value3 bool     `toml:"value3"` // ブール
			Value4 []string `toml:"value4"` // リスト
		}

		AuthorSection struct {
			Name string // タグ指定なし
		}

		Person struct {
			Name string
			Age  int
		}

		Root struct {
			Values  ValuesSection
			Author  AuthorSection
			Persons []Person
		}
	)

	var (
		v   Root
		err error
	)

	err = toml.Unmarshal([]byte(src), &v)
	if err != nil {
		return err
	}

	output.Stdoutf("[Unmarshal]", "%+v\n", v)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: toml_unmarshal

	   [Name] "toml_unmarshal"
	   [Unmarshal]          {Values:{Value1:999 Value2:hello world Value3:false Value4:[one two three]} Author:{Name:devlights} Persons:[{Name:one Age:30} {Name:two Age:0} {Name:three Age:99}]}


	   [Elapsed] 111.32µs
	*/

}

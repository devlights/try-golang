package tomlop

import (
	"github.com/devlights/gomy/output"
	"github.com/pelletier/go-toml/v2"
)

// Marshal は、TOMLファイルの内容を書き込むサンプルです.
//
// # REFERENCES
//   - https://www.meetgor.com/golang-config-file-read/
//   - https://github.com/pelletier/go-toml
//   - https://ja.wikipedia.org/wiki/TOML
func Marshal() error {
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
			Name string `toml:"name"`
			Age  int    `toml:"age,omitempty"`
		}

		Root struct {
			Values  ValuesSection
			Author  AuthorSection
			Persons []Person
		}
	)

	var (
		valuesSection = ValuesSection{
			Value1: 999,
			Value2: "hello世界",
			Value3: false,
			Value4: []string{"Go", "C#", "Python"},
		}
		authorSection = AuthorSection{
			Name: "devlights",
		}
		persons = []Person{
			{Name: "one", Age: 30},
			{Name: "two"},
			{Name: "three", Age: 99},
		}
		root = Root{
			Values:  valuesSection,
			Author:  authorSection,
			Persons: persons,
		}
	)

	var (
		serialized []byte
		err        error
	)

	serialized, err = toml.Marshal(&root)
	if err != nil {
		return err
	}

	output.Stdoutf("[Marshal]", "\n%s\n", string(serialized))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: toml_marshal

	   [Name] "toml_marshal"
	   [Marshal]
	   [Values]
	   value1 = 999
	   value2 = 'hello世界'
	   value3 = false
	   value4 = ['Go', 'C#', 'Python']

	   [Author]
	   Name = 'devlights'

	   [[Persons]]
	   name = 'one'
	   age = 30

	   [[Persons]]
	   name = 'two'

	   [[Persons]]
	   name = 'three'
	   age = 99



	   [Elapsed] 60.74µs
	*/

}

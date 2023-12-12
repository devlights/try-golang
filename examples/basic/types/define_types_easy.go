package types

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// 型定義
type (
	Id   string
	Name string

	User struct {
		Id   Id
		Name Name
	}
)

func (me User) String() string {
	return fmt.Sprintf("id:%s\tname:%s", me.Id, me.Name)
}

// DefineTypesEasy -- Goでは型定義が簡単であるということを示すサンプルです.
func DefineTypesEasy() error {
	var (
		id   = Id("0001")
		name = Name("golang")
	)

	// よくあるパターンが 関数定義 が以下のようになっていて
	//   func newUser(id, name string)
	// 渡す引数を間違えてしまうというパターン。
	//   ex: newUser(name, id)
	// 下記のように型定義されていれば間違えることもない (コンパイルエラーになるため)
	//   func newUser(id Id, name Name)
	// Go では型定義がとても簡単に出来るので、これがやりやすい
	user, err := newUser(id, name)
	if err != nil {
		return err
	}

	output.Stdoutl("[result]", user)

	return nil
}

func newUser(id Id, name Name) (User, error) {
	return User{Id: id, Name: name}, nil
}

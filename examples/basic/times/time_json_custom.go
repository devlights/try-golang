package times

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/devlights/gomy/errs"
	"github.com/devlights/gomy/output"
	"github.com/devlights/gomy/times"
)

type jsonTimeCustom struct {
	TimeStamp time.Time `json:"timestamp"`
}

// Implements
var (
	_ fmt.Stringer     = (*jsonTimeCustom)(nil)
	_ json.Marshaler   = (*jsonTimeCustom)(nil)
	_ json.Unmarshaler = (*jsonTimeCustom)(nil)
)

func (me *jsonTimeCustom) String() string {
	return times.Formatter(me.TimeStamp).Format("yyyy-MM-dd hh:mm:ss")
}

func (me *jsonTimeCustom) MarshalJSON() ([]byte, error) {
	// json の場合、ダブルクォートで囲っていないとエラーになるので注意
	return []byte(fmt.Sprintf("\"%s\"", me)), nil
}

func (me *jsonTimeCustom) UnmarshalJSON(b []byte) error {
	// json の場合、値はダブルクォートで囲まれた状態となっているので注意
	t, err := times.Parser(string(b)).Parse("\"yyyy-MM-dd hh:mm:ss\"")
	if err != nil {
		return err
	}

	me.TimeStamp = t

	return nil
}

// TimeJsonCustom -- time.Time をカスタム JSON エンコード・デコードで利用するサンプルです.
//
// REFERENCES
//   - https://zenn.dev/hsaki/articles/go-time-cheatsheet#jst(%E6%97%A5%E6%9C%AC%E6%A8%99%E6%BA%96%E6%99%82)%E3%82%92%E6%89%B1%E3%81%86%E5%A0%B4%E5%90%88
func TimeJsonCustom() error {
	var (
		loc = errs.Forgot(time.LoadLocation("Asia/Tokyo"))
		t   = time.Now().In(loc)
		v1  = jsonTimeCustom{t}
		v2  = jsonTimeCustom{}
	)

	b, err := json.Marshal(&v1)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &v2)
	if err != nil {
		return err
	}

	output.Stdoutl("[v1]", &v1)
	output.Stdoutl("[v2]", &v2)

	return nil
}

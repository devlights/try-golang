package times

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/devlights/gomy/output"
	"github.com/devlights/gomy/times"
)

type jsonTime struct {
	Timestamp time.Time `json:"timestamp"`
}

var _ fmt.Stringer = (*jsonTime)(nil)

func (me *jsonTime) String() string {
	if me == nil {
		return ""
	}

	return times.Formatter(me.Timestamp).Format("yyyy-MM-dd HH:mm:ss")
}

// TimeJson -- time.Time を json 形式で扱う場合のサンプルです。
//
// REFERENCES
//   - https://zenn.dev/hsaki/articles/go-time-cheatsheet#time.time%E5%9E%8B--%3E-json%E6%96%87%E5%AD%97%E5%88%97
func TimeJson() error {
	var (
		encode = func() ([]byte, error) {
			t := time.Now()

			b, err := json.Marshal(jsonTime{t})
			if err != nil {
				return nil, err
			}

			return b, err
		}
		decode = func(b []byte) (*jsonTime, error) {
			var v jsonTime

			err := json.Unmarshal(b, &v)
			if err != nil {
				return nil, err
			}

			return &v, err
		}
	)

	b, err := encode()
	if err != nil {
		return err
	}

	v, err := decode(b)
	if err != nil {
		return err
	}

	output.Stdoutf("[encode]", "%s\n", b)
	output.Stdoutf("[decode]", "%s\n", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_json

	   [Name] "time_json"
	   [encode]             {"timestamp":"2024-03-27T06:02:25.537803295Z"}
	   [decode]             2024-03-27 HH:02:25


	   [Elapsed] 147.55µs
	*/

}

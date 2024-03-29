package types

import (
	"encoding/json"
	"fmt"
	_ "strconv"
	"time"
)

type (
	// YyyyMmDd は、 YYYY/MM/DD 形式で json.Marshal/json.Unmarshal するために利用できる構造体です.
	YyyyMmDd struct {
		time.Time
	}
)

var (
	_ fmt.Stringer     = (*YyyyMmDd)(nil)
	_ json.Marshaler   = (*YyyyMmDd)(nil)
	_ json.Unmarshaler = (*YyyyMmDd)(nil)
)

func (me YyyyMmDd) String() string {
	return me.string()
}

func (me YyyyMmDd) MarshalJSON() ([]byte, error) {
	return []byte(me.jsonString()), nil
}

func (me *YyyyMmDd) UnmarshalJSON(b []byte) error {
	var (
		s   = string(b)
		err error
	)

	if s == "null" {
		return nil
	}

	// https://stackoverflow.com/questions/16846553/how-to-unmarshal-an-escaped-json-string
	// if s, err = strconv.Unquote(s); err != nil {
	// 	return err
	// }

	var (
		t time.Time
	)

	// https://essential-go.programming-books.io/custom-json-marshaling-468765d144a34e87b913c7674e66c3a4#12ec5c24-ae51-4341-95ae-a5a0a81ed746
	if t, err = time.Parse(`"2006/01/02"`, s); err != nil {
		return err
	}

	me.Time = t

	return nil
}

func (me YyyyMmDd) string() string {
	return me.Time.Format("2006/01/02")
}

func (me YyyyMmDd) jsonString() string {
	return fmt.Sprintf(`"%s"`, me.string())
}

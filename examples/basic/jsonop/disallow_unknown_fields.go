package jsonop

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/devlights/gomy/output"
)

// DisallowUnknownFields は、*Decoder.DisallowUnknownFields のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/encoding/json@go1.21.6#Decoder.DisallowUnknownFields
func DisallowUnknownFields() error {
	const (
		jsonValue = `{"id": 1, "name": "Test1", "age": 99}`
	)

	type Val struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var (
		stream = strings.NewReader(jsonValue)
		dec    = json.NewDecoder(stream)
	)

	//
	// 普通にデコード
	//
	// *json.Decoder.DisallowUnknownFields() を呼んでいないので
	// 存在しないJSONフィールドがあってもエラーにはならない。
	//
	var (
		v   Val
		err error
	)

	err = dec.Decode(&v)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	output.Stdoutf("[Normal]", "%v\n", v)

	//
	// 不明なフィールドは許可しないよう設定
	//
	// この場合の「不明なフィールド」というのは
	// "JSON側に存在しているフィールドが、受け側の構造体に存在しない場合" という意味。
	// 構造体側に存在しているフィールドが、JSON側に存在しないのはエラーにならないので注意。
	//
	// 今度は、ageというJSONフィールドに対応する構造体フィールドが存在しないのでエラーとなる。
	//
	stream = strings.NewReader(jsonValue)
	dec = json.NewDecoder(stream)

	dec.DisallowUnknownFields()

	err = dec.Decode(&v)
	if err != nil && !errors.Is(err, io.EOF) {
		output.Stdoutf("[DisallowUnknownFields]", "%v(%T)", err, err)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_disallow_unknown_fields

	   [Name] "json_disallow_unknown_fields"
	   [Normal]             {1 Test1}
	   [DisallowUnknownFields] json: unknown field "age"(*errors.errorString)

	   [Elapsed] 158.11µs
	*/
}

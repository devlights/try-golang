package csvop

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

// FieldsPerRecord は、csv.Reader.FieldsPerRecordのサンプルです。
//
// csv.Reader.FieldsPerRecord は設定する値によって挙動が変わる。
//
//   - 「0より大きな正の値」を設定すると、その列数でない場合 *csv.ParseError が発生する
//   - 「0」を指定すると、最初の行の列数を基準として解析し、その列数でない場合 *csv.ParseError が発生する
//   - 「負の値」 を設定すると列が不揃いでもエラーにならない
//
// REFERENCES:
//   - https://pkg.go.dev/encoding/csv@go1.25.3#Reader.FieldsPerRecord
func FieldsPerRecord() error {
	var (
		data []byte
	)
	data = fmt.Appendln(data, "hello,world")
	data = fmt.Appendln(data, "world,hello")
	data = fmt.Appendln(data, "HELLO,WORLD,999") // ここだけ列が不揃い

	//
	// FieldsPerRecordの値が「０より大きな正の値」
	// (指定した列数で無い場合エラーとなる)
	//
	{
		var (
			buf    = bytes.NewBuffer(data)
			reader = csv.NewReader(buf)
			record []string
			err    error
		)
		reader.FieldsPerRecord = 2

		for i := 0; ; i++ {
			if record, err = reader.Read(); err != nil {
				fmt.Printf("[０より大きな正の値] [%d] %T: %s\n", i, err, err)
				break
			}

			fmt.Printf("[０より大きな正の値] [%d] %v\n", i, record)
		}
	}

	//
	// FieldsPerRecordの値が「０」
	// (先頭レコードの列数を基準として処理する)
	//
	{
		var (
			buf    = bytes.NewBuffer(data)
			reader = csv.NewReader(buf)
			record []string
			err    error
		)
		reader.FieldsPerRecord = 0

		for i := 0; ; i++ {
			if record, err = reader.Read(); err != nil {
				fmt.Printf("[０　　　　　　　　] [%d] %T: %s\n", i, err, err)
				break
			}

			fmt.Printf("[０　　　　　　　　] [%d] %v\n", i, record)
		}
	}

	//
	// FieldsPerRecordの値が「負の値」
	// (列数が不揃いでもエラーとならない)
	//
	{
		var (
			buf    = bytes.NewBuffer(data)
			reader = csv.NewReader(buf)
			record []string
			err    error
		)
		reader.FieldsPerRecord = -1

		for i := 0; ; i++ {
			if record, err = reader.Read(); err != nil {
				fmt.Printf("[負の値　　　　　　] [%d] %T: %s\n", i, err, err)
				break
			}

			fmt.Printf("[負の値　　　　　　] [%d] %v\n", i, record)
		}
	}

	return nil

	/*
		$ task
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: csv_fields

		[Name] "csv_fieldsperrecord"
		[０より大きな正の値] [0] [hello world]
		[０より大きな正の値] [1] [world hello]
		[０より大きな正の値] [2] *csv.ParseError: record on line 3: wrong number of fields
		[０　　　　　　　　] [0] [hello world]
		[０　　　　　　　　] [1] [world hello]
		[０　　　　　　　　] [2] *csv.ParseError: record on line 3: wrong number of fields
		[負の値　　　　　　] [0] [hello world]
		[負の値　　　　　　] [1] [world hello]
		[負の値　　　　　　] [2] [HELLO WORLD 999]
		[負の値　　　　　　] [3] *errors.errorString: EOF


		[Elapsed] 352.396µs
	*/
}

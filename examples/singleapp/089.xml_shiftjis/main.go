// XML宣言にてencodingの指定がUTF-8ではない場合のXMLデコードのサンプルです.
//
// # REFERENCES
//
//   - https://stackoverflow.com/questions/54915307/error-unmarshaling-a-simple-xml-in-golang
//   - https://qiita.com/bamchoh/items/a4c64ace78200bf0fa6e
//   - https://pkg.go.dev/encoding/xml@go1.19.2#Decoder
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"

	"github.com/devlights/gomy/fileio"
	"github.com/devlights/gomy/fileio/jp"
)

type (
	xmlData struct {
		XMLName xml.Name `xml:"data"`
		Hello   string   `xml:"hello"`
		World   string   `xml:"world"`
	}
)

func (me xmlData) String() string {
	return fmt.Sprintf("hello=%s\tworld=%s", me.Hello, me.World)
}

func fail() error {
	var (
		r      io.Reader
		closer func() error
		err    error
	)

	r, closer, err = fileio.OpenRead("sample.xml", jp.ShiftJis)
	if err != nil {
		return err
	}
	defer closer()

	var (
		data    xmlData
		decoder = xml.NewDecoder(r)
	)

	// encoding="shift-jis" なデータをそのままUnmarshalするとエラーになる
	//
	// 以下のランタイムエラーが出る
	//   xml: encoding "shift-jis" declared but Decoder.CharsetReader is nil
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("[fail] %v\n", err)
		return nil
	}

	fmt.Printf("[fail] %v\n", data)

	return nil
}

func succ() error {
	var (
		r      io.Reader
		closer func() error
		err    error
	)

	r, closer, err = fileio.OpenRead("sample.xml", jp.ShiftJis)
	if err != nil {
		return err
	}
	defer closer()

	var (
		data    xmlData
		decoder = xml.NewDecoder(r)
	)

	// encoding="shift-jis" なデータをそのままUnmarshalするとエラーになる
	//
	// XML宣言にてUTF-8以外のエンコーディングが指定されている場合
	// CharsetReaderが呼び出されるため、設定する必要がある.
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		// 既に shift-jis で読み出せる io.Reader なので、そのまま返す.
		// そうではない場合は、ここでラップして返す.
		return input, nil
	}

	err = decoder.Decode(&data)
	if err != nil {
		return err
	}

	fmt.Printf("[succ] %v\n", data)

	return nil
}

func run() error {
	var (
		err error
	)

	err = fail()
	if err != nil {
		return err
	}

	err = succ()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

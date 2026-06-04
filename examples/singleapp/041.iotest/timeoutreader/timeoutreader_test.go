package timeoutreader

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

// TestReadAllAtOnce は、io.ReadAll()を利用してデータを読み取るテストケースです。
// テストで利用している io.Reader は、iotest.TimeoutReader()から生成しています。
//
// iotest.TimeoutReader() は、**２回目のみ**タイムアウトで失敗する
// io.Readerを作成してくれます。なので、通信処理などをテストする際にとても便利です。
// ２回目以降は普通にデータが読み取れるようになっています。
//
// このテストケースでは、読み取りに io.ReadAll() を利用しているため、２回目のエラーで
// 処理が返ってしまい、テストケースがFailとなります。io.ReadAll()には、リトライ処理などは
// 実装されていないため、タイムアウトが発生する可能性がある場面では利用しない方が良いということになります。
//
// # REFERENCES
//   - https://pkg.go.dev/testing/iotest@go1.23.0#TimeoutReader
func TestReadAllAtOnce(t *testing.T) {
	var (
		data   = make([]byte, 1<<10)
		buf    = make([]byte, len(data))
		reader io.Reader
		err    error
	)

	randomString(data)
	reader = iotest.TimeoutReader(bytes.NewReader(data))

	err = readAllAtOnce(reader, buf)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, buf) {
		t.Fatalf("[want] equal\t[got] not equal")
	}
}

// TestReadWithRetry は、リトライ処理を考慮した読み取り処理を利用してデータを読み取るテストケースです。
// テストで利用している io.Reader は、iotest.TimeoutReader()から生成しています。
//
// iotest.TimeoutReader() は、**２回目のみ**タイムアウトで失敗する
// io.Readerを作成してくれます。なので、通信処理などをテストする際にとても便利です。
// ２回目以降は普通にデータが読み取れるようになっています。
//
// このテストケースでは、読み取り中にタイムアウトを含むエラーが発生しても、指定回数リトライして
// 再試行するようになっているため、テストケースが通ります。
//
// # REFERENCES
//   - https://pkg.go.dev/testing/iotest@go1.23.0#TimeoutReader
func TestReadWithRetry(t *testing.T) {
	var (
		data   = make([]byte, 1<<10)
		buf    = make([]byte, len(data))
		reader io.Reader
		err    error
	)

	randomString(data)
	reader = iotest.TimeoutReader(bytes.NewReader(data))

	err = readWithRetry(reader, buf, 3)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, buf) {
		t.Fatalf("[want] equal\t[got] not equal")
	}
}

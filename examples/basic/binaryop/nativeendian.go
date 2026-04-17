package binaryop

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
	"unsafe"
)

// NativeEndian は、binary.NativeEndian のサンプルです。
//
// encoding/binary には、バイトオーダーを扱う３つの変数が定義されている。
//
//   - binary.BigEndian    : 最上位バイトから順に格納（ネットワークバイトオーダーとも呼ばれる）
//   - binary.LittleEndian : 最下位バイトから順に格納（x86/x86-64系CPUの標準）
//   - binary.NativeEndian : 実行環境のCPUに合わせたバイトオーダー
//
// 昨今のデスクトップ・サーバー向けCPU（x86/x86-64）はリトルエンディアンが標準であり、
// ARMもiOS/Androidを含む多くの環境でリトルエンディアンモードで動作する。
// そのため、NativeEndianはほぼリトルエンディアンと同義になるケースが多い。
//
// NativeEndianの主な用途は、ローカルIPC（共有メモリ・パイプ等）のように
// 送受信が同一ホスト内で完結し、エンディアンを意識せずに記述したい場面で有用となる。
// ネットワーク通信ではBigEndian（ネットワークバイトオーダー）を明示するのが原則。
//
// # REFERENCES
//   - https://pkg.go.dev/encoding/binary@go1.26.2#pkg-variables
func NativeEndian() error {
	var (
		v  uint32 = 0x01020304
		sz        = int(unsafe.Sizeof(v))
		be        = make([]byte, sz) // big-endian
		le        = make([]byte, sz) // little-endian
		ne        = make([]byte, sz) // native-endian
	)
	binary.BigEndian.PutUint32(be, v)
	binary.LittleEndian.PutUint32(le, v)
	binary.NativeEndian.PutUint32(ne, v)

	fmt.Printf("[Original     ] 0x%08X\n", v)
	fmt.Printf("[Big    Endian] %v\n", be)
	fmt.Printf("[Little Endian] %v\n", le)
	fmt.Printf("[Native Endian] %v\n", ne)
	fmt.Println("-----------------------")

	//
	// bianry.NativeEndianで読み取るサンプル
	// わざわざソケット用意するのは面倒なので os.Pipe() で代用
	//
	var (
		pr, pw *os.File
		err    error
	)
	if pr, pw, err = os.Pipe(); err != nil {
		return nil
	}

	// SEND SIDE
	const (
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)
	var (
		rnd = func(n int) []byte {
			b := make([]byte, n)
			for i := range b {
				b[i] = charset[rand.IntN(len(charset))]
			}
			return b
		}
		errCh = make(chan error, 1)
	)
	go func() {
		defer pw.Close()

		var (
			buf    = new(bytes.Buffer)
			data   = rnd(10 + rand.IntN(21))
			length = uint32(len(data))
		)
		// ローカルパイプ上なので、エンディアン意識せずにバイナリを送りたい
		binary.Write(buf, binary.NativeEndian, length)
		binary.Write(buf, binary.NativeEndian, data)

		if _, err = pw.Write(buf.Bytes()); err != nil {
			errCh <- err
		}
	}()

	// RECV SIDE
	var (
		header = make([]byte, 4)
	)
	if _, err = pr.Read(header); err != nil {
		return err
	}

	var (
		// ローカルパイプ上なので、エンディアン意識せずにバイナリを読み取りたい
		length  = binary.NativeEndian.Uint32(header)
		payload = make([]byte, length)
	)
	if _, err = pr.Read(payload); err != nil {
		return err
	}

	fmt.Printf("length=%d, payload=%s\n", length, payload)

	select {
	case err = <-errCh:
		return err
	default:
	}

	return nil
}

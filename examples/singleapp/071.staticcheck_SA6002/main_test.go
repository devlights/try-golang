package main

import (
	"sync"
	"testing"
)

const (
	BUF_LEN = 1e6
)

// useBuffer は、 []byte を利用したロジックです.
func useBuffer(buf []byte) {
	clear(buf)
	buf = buf[:0]

	for i := 0; i < BUF_LEN; i++ {
		buf = append(buf, 'a')
	}

	// Use buf...
}

// usePoolWithPointer は、sync.Pool を利用したロジックです.
//
// プールの中には スライスのポインタ が格納されています。
// これは staticcheck の SA6002 の指摘に従ったやり方です。
//
// # REFERENCES
//   - https://staticcheck.dev/docs/checks#SA6002
//   - https://github.com/dominikh/go-tools/issues/1336
//   - https://github.com/dominikh/go-tools/issues/302
func usePoolWithSlicePointer(p *sync.Pool) {
	bufPtr := p.Get().(*[]byte)
	buf := *bufPtr

	clear(buf)
	buf = buf[:0]

	for i := 0; i < BUF_LEN; i++ {
		buf = append(buf, 'a')
	}

	defer func() {
		// bufのスライスヘッダが編集によって変わってしまっている
		// 可能性（サイズや配列へのポインタが変更された場合に備えて）を考慮して
		// p.Put(&buf) とするのでは無く、元々プールに存在しているポインタである
		// bufPtr に上書きしてからプールに戻す。
		//
		// REF: https://github.com/dominikh/go-tools/issues/1336#issuecomment-1331206290
		*bufPtr = buf
		p.Put(bufPtr)
	}()

	// Use buf...
}

// usePoolWithSliceDirect は、sync.Pool を利用したロジックです.
//
// プールの中には スライス が格納されています。
// このコードは staticcheck だと SA6002 として警告されます。
//
// ですが、余分な割当が行われるという点の警告であるので
// 別段やってはいけない処理の書き方ではありません。
//
// 実際にベンチマークを取ると SA6002 に従った書き方よりも
// こちらの方が速度は速くなります。ただし、余計な割当が入る可能性がある。
//
// REF: https://github.com/dominikh/go-tools/issues/1336
func usePoolWithSliceDirect(p *sync.Pool) {
	buf := p.Get().([]byte)

	clear(buf)
	buf = buf[:0]

	for i := 0; i < BUF_LEN; i++ {
		buf = append(buf, 'a')
	}

	defer p.Put(buf) //lint:ignore SA6002 非ポインタを意図的に利用しているので問題無し

	// Use buf...
}

func BenchmarkStaticCheckSA6002(b *testing.B) {
	b.Run("alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			useBuffer(nil)
		}
	})

	b.Run("buffer", func(b *testing.B) {
		buf := make([]byte, BUF_LEN)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			useBuffer(buf)
		}
	})

	b.Run("pool-sa6002-ok", func(b *testing.B) {
		pool := sync.Pool{
			New: func() any {
				buf := make([]byte, BUF_LEN)
				return &buf
			},
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			usePoolWithSlicePointer(&pool)
		}
	})

	b.Run("pool-sa6002-ng", func(b *testing.B) {
		pool := sync.Pool{
			New: func() any {
				return make([]byte, BUF_LEN)
			},
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			usePoolWithSliceDirect(&pool)
		}
	})
}

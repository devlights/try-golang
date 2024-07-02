package bufferop

import (
	"bytes"

	"github.com/devlights/gomy/output"
)

// AvailableBuffer は、Go 1.21 で追加された Buffer.AvailableBuffer() についてのサンプルです。
//
// > AvailableBuffer returns an empty buffer with b.Available() capacity.
// This buffer is intended to be appended to and passed to an immediately succeeding Buffer.Write call.
// The buffer is only valid until the next write operation on b.
//
// > (AvailableBufferは、b.Available()の容量を持つ空のバッファを返します。
// このバッファは、直後のBuffer.Write呼び出しに追加され、渡されることを意図しています。
// このバッファは、bに対する次の書き込み操作が行われるまで有効です。)
//
// # REFERENCES
//   - https://pkg.go.dev/bytes@go1.22.4#Buffer.AvailableBuffer
func AvailableBuffer() error {
	var (
		buf   bytes.Buffer
		data  []byte
		logfn = func(buf *bytes.Buffer, data []byte) {
			output.Stdoutl("[Available      ]", buf.Available())
			output.Stdoutf("[AvailableBuffer]", "data=%v\tlen=%d\tcap=%d\n", data, len(data), cap(data))
			output.Stdoutl("[Bytes Length   ]", len(buf.Bytes()))
			output.StdoutHr()
		}
	)

	buf.Grow(1 << 11)

	data = buf.AvailableBuffer()
	{
		logfn(&buf, data)

		data = append(data, []byte("helloworld")...)
		_, _ = buf.Write(data)
	}

	data = buf.AvailableBuffer()
	{
		logfn(&buf, data)

		var b [1 << 10]byte
		data = append(data, b[:]...)
		_, _ = buf.Write(data)
	}

	data = buf.AvailableBuffer()
	{
		logfn(&buf, data)
	}

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: bufferop_available_buffer

		[Name] "bufferop_available_buffer"
		[Available      ]    2048
		[AvailableBuffer]    data=[]    len=0   cap=2048
		[Bytes Length   ]    0
		--------------------------------------------------
		[Available      ]    2038
		[AvailableBuffer]    data=[]    len=0   cap=2038
		[Bytes Length   ]    10
		--------------------------------------------------
		[Available      ]    1014
		[AvailableBuffer]    data=[]    len=0   cap=1014
		[Bytes Length   ]    1034
		--------------------------------------------------


		[Elapsed] 105.66µs
	*/

}

# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                  | example name                | note                                                                 |
| --------------------- | --------------------------- | -------------------------------------------------------------------- |
| zerovalue.go          | bufferop_zero_value         | bytes.Buffer を ゼロ値 で利用した場合のサンプルです                  |
| frombytes.go          | bufferop_from_bytes         | bytes.Buffer を バイト列 から生成するサンプルです                    |
| fromstring.go         | bufferop_from_string        | bytes.Buffer を 文字列 から生成するサンプルです                      |
| use_as_reader.go      | bufferop_use_as_reader      | bytes.Buffer を io.Reader として利用するサンプルです                 |
| use_as_writer.go      | bufferop_use_as_writer      | bytes.Buffer を io.Writer として利用するサンプルです                 |
| availablebuffer.go    | bufferop_available_buffer   | Go 1.21 で追加された Buffer.AvailableBuffer() についてのサンプルです |
| to_readwritecloser.go | bufferop_to_readwritecloser | bytes.Buffer を io.ReadWriteCloser に変換するサンプルです            |

# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                  | example name            | note                                                                                     |
| --------------------- | ----------------------- | ---------------------------------------------------------------------------------------- |
| limitread.go          | ioop_limit_read         | io.LimitedReader のサンプルです.                                                         |
| onebyteread.go        | ioop_onebyte_read       | １バイトずつ読み出す io.LimitedReader のサンプルです.                                    |
| multiwrite.go         | ioop_multiwrite         | io.MultiWriterを利用してgzip圧縮しながらCRCチェックサムも算出するサンプルです.           |
| multiread.go          | ioop_multiread          | io.MultiReaderを利用して複数のファイルを一気に読み込むサンプルです。                     |
| teeread.go            | ioop_tee_read           | io.TeeReader を利用したサンプルです。                                                    |
| sectionread.go        | ioop_section_read       | io.SectionReader を利用したサンプルです。                                                |
| offsetwrite.go        | ioop_offset_write       | io.OffsetWriter を利用したサンプルです。                                                 |
| multi_offsetwriter.go | ioop_offset_write_multi | io.OffsetWriter を非同期で複数実行し、それぞれ異なるオフセット位置に書き込むサンプルです |

# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                                 | example name                | note                                                                 |
|--------------------------------------|-----------------------------|----------------------------------------------------------------------|
| cmpwaitgroup/error_with_errgroup.go  | errgrp_error_with_errgroup  | 拡張ライブラリ golang.org/x/sync/errgroup でエラー情報を呼び元に伝播させるサンプルです            |
| cmpwaitgroup/error_with_waitgroup.go | errgrp_error_with_waitgroup | 標準ライブラリ sync.WaitGroup でエラー情報を呼び元に伝播させるサンプルです                        |
| pipeline/with_pipeline               | errgrp_with_pipeline        | 拡張ライブラリ golang.org/x/sync/errgroup でパイプライン処理を行っているサンプルです             |
| withcontext/with_context.go          | errgrp_with_context         | 拡張ライブラリ golang.org/x/sync/errgroup で ctx.Context を含めた利用方法についてのサンプルです |

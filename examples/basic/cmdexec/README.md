# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                       | example name                    | note                                                                       |
| -------------------------- | ------------------------------- | -------------------------------------------------------------------------- |
| oneshot.go                 | cmdexec_oneshot                 | コマンドを一発実行して結果を取得するサンプルです                           |
| oneshotwithstderr.go       | cmdexec_oneshot_with_stderr     | コマンドを一発実行して結果を取得するサンプルです。(標準エラー出力も含む)   |
| stdinouterr.go             | cmdexec_stdinouterr             | 標準入力・標準出力・標準エラー出力を指定してコマンドを実行するサンプルです |
| withcontext.go             | cmdexec_withcontext             | コマンドを context.Context 付きで実行するサンプルです                      |
| pipe.go                    | cmdexec_pipe                    | (*Cmd).StdinPipe,StdoutPipe,StderrPipeのサンプルです                       |
| multi_command_with_pipe.go | cmdexec_multi_command_with_pipe | 複数の (*exec.Cmd) をパイプストリームで繋いで実行するサンプルです          |
| withenv.go                 | cmdexec_env                     | *exec.Cmd 実行時に追加の環境変数を指定するサンプルです                     |
| withdir.go                 | cmdexec_dir                     | *exec.Cmd 実行時にワーキングディレクトリを指定するサンプルです             |
| withslice.go               | cmdexec_slice                   | *exec.Cmd 実行時にスライスの値をコマンドの引数で指定するサンプルです       |

# サンプルリスト

このディレクトリには以下のサンプルがあります。

|file|example name|note|
|----|------------|----|
|oneshot.go|cmdexec\_oneshot|コマンドを一発実行して結果を取得するサンプルです|
|oneshotwithstderr.go|cmdexec\_oneshot\_with\_stderr|コマンドを一発実行して結果を取得するサンプルです。(標準エラー出力も含む)|
|stdinouterr.go|cmdexec\_stdinouterr|標準入力・標準出力・標準エラー出力を指定してコマンドを実行するサンプルです|
|withcontext.go|cmdexec\_withcontext|コマンドを context.Context 付きで実行するサンプルです|
|pipe.go|cmdexec\_pipe|(*Cmd).StdinPipe,StdoutPipe,StderrPipeのサンプルです|

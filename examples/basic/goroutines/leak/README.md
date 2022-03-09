# サンプルリスト

このディレクトリには以下のサンプルがあります。

|file|example name|note|
|----|------------|----|
|forgotten\_sender.go|goroutines\_leak\_forgotten\_sender|チャネルの送信側を忘れることにより発生するgoroutineリークのサンプルです|
|forgotten\_receiver.go|goroutines\_leak\_forgotten\_receiver|チャネルの受信側を忘れることにより発生するgoroutineリークのサンプルです|
|abandoned\_sender.go|goroutines\_leak\_abandoned\_sender|処理のタイミングによって受信側がいなくなり、送信側が放棄されてしまうgoroutineリークのサンプルです|
|abandoned\_receiver.go|goroutines\_leak\_abandoned\_receiver|処理のタイミングによって送信側がいなくなり、受信側が放棄されてしまうgoroutineリークのサンプルです|
|sender\_after\_error\_check.go|goroutines\_leak\_sender\_after\_error\_check|処理結果によって送信側がいなくなり、受信側が放棄されてしまうgoroutineリークのサンプルです|
# サンプルリスト

このディレクトリには以下のサンプルがあります。

|file|example name|note|
|----|------------|----|
|nonstop.go|goroutines\_nonstop|ゴルーチンを待ち合わせ無しで走らせるサンプルです|
|withdonechannel.go|goroutines\_with\_done\_channel|doneチャネルを用いて待ち合わせを行うサンプルです|
|withwaitgroup.go|goroutines\_with\_waitgroup|sync.WaitGroupを用いて待ち合わせを行うパターンです|
|withcontextcancel.go|goroutines\_with\_context\_cancel|context.Contextを用いて待ち合わせを行うサンプルです|
|withcontexttimeout.go|goroutines\_with\_context\_timeout|context.Contextを用いてタイムアウト付きで待ち合わせを行うサンプルです|
|selectnilchan1.go|goroutines\_select\_nil\_chan\_1|select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (1)|
|selectnilchan2.go|goroutines\_select\_nil\_chan\_2|select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (2)|
|usingchansemaphore.go|goroutines_using_chan_semaphore|チャネルでセマフォの動作を行わせるサンプルです|
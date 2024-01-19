# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                         | example name                         | note                                                                                                    |
| ---------------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------- |
| nonstop.go                   | goroutines_nonstop                   | ゴルーチンを待ち合わせ無しで走らせるサンプルです                                                        |
| withdonechannel.go           | goroutines_with_done_channel         | doneチャネルを用いて待ち合わせを行うサンプルです                                                        |
| withwaitgroup.go             | goroutines_with_waitgroup            | sync.WaitGroupを用いて待ち合わせを行うパターンです                                                      |
| withcontextcancel.go         | goroutines_with_context_cancel       | context.Contextを用いて待ち合わせを行うサンプルです                                                     |
| withcontexttimeout.go        | goroutines_with_context_timeout      | context.Contextを用いてタイムアウト付きで待ち合わせを行うサンプルです                                   |
| withcontextdeadline.go       | goroutines_with_context_deadline     | context.Context::WithDeadline を使ったサンプルです                                                      |
| selectnilchan1.go            | goroutines_select_nil_chan_1         | select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (1) |
| selectnilchan2.go            | goroutines_select_nil_chan_2         | select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (2) |
| usingchansemaphore.go        | goroutines_using_chan_semaphore      | チャネルでセマフォの動作を行わせるサンプルです                                                          |
| usingmutex.go                | goroutines_using_mutex               | sync.Mutex を利用したサンプルです                                                                       |
| workerpool.go                | goroutines_workerpool                | Worker Pool パターンのサンプルです                                                                      |
| context_and_timeafterfunc.go | goroutines_context_and_timeafterfunc | ContextAndTimeAfterFunc は、Context と time.AfterFunc でキャンセルするサンプルです                      |

# これは何？

Go1.24で実験的機能として追加され、Go1.25で一般提供となった [testing.synctest](https://pkg.go.dev/testing/synctest) パッケージについてのサンプルが配置されています。

## bubble について

**bubble** とは、テスト中に生成したgoroutineやtime.Timerなどのリソースを外部と完全に隔離した実行環境を意味します。

### bubbleの特徴

#### 隔離された実行環境

synctest.Test()で作成されるbubble内で起動したgoroutineや生成したチャネル、タイマーなどは、bubble外から操作できません。たとえば、bubble内で作ったチャネルをbubble外から送受信しようとするとpanicになります。

#### 専用のfake時間

bubble内ではtimeパッケージがfake clock（模擬時計）を使います。時間はbubble内のgoroutineがすべて**durably blocked（耐久ブロック）**状態になったときのみ進みます。つまり、計算処理中は時間が止まり、SleepやTimerの待ちが発生したときのみ時間が進みます。

#### durably blocked（耐久ブロック）

goroutineが「耐久ブロック」状態とは、そのgoroutineがbubble内の他のgoroutineからのみ解放される状態を指します。たとえば、bubble内のチャネルに対する送受信、sync.Cond.Wait、sync.WaitGroup.Wait、time.Sleepなどが該当します。
逆に、ファイルI/OやネットワークI/Oなど、bubble外の要因でブロックされる場合は「耐久ブロック」とはみなされません。

#### テストの安定性と高速化

bubbleを使うことで、goroutineの競合やタイミング依存によるテストの不安定さ（フレーキー）を防ぎつつ、実際には待たずに一瞬でテストが終わるため、高速かつ確実な非同期テストが可能になります。

### bubbleの注意点

#### I/O系のブロックは対象外

ネットワークやファイルI/Oなど、カーネルや外部システムに依存するブロックは「耐久ブロック」とみなされず、bubbleの管理外です。そのため、bubble内でネットワーク通信をテストしたい場合は、fakeネットワーク（例：net.Pipe()）を使う必要があります。

#### bubbleのライフサイクル

bubbleのルートgoroutine（Testに渡した関数）が終了すると、それ以降はfake時間も進まなくなり、bubble内のgoroutineがSleepやTimerでブロックしているとデッドロックでテストが失敗します。


## 参考情報

- [synctest package](https://pkg.go.dev/testing/synctest)
- [Testing Time (and other asynchronicities)](https://go.dev/blog/testing-time)
- [Go synctest: Solving Flaky Tests](https://victoriametrics.com/blog/go-synctest/)
- [Gist of Go: Concurrency testing](https://antonz.org/go-concurrency/testing/)
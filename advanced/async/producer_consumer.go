package async

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Flags() &^ log.LstdFlags)
}

// Main, Producer, Consumer, RemainCollector が利用するデータ
type (
	item  int
	empty struct{}
)

// Main, Producer, Consumer, RemainCollector の間で利用されるチャネル達
type (
	done      <-chan empty
	itemCh    <-chan item
	terminate <-chan empty
)

// 各役割ごとに決められたインターバル
var (
	producerInterval     = 500 * time.Millisecond
	consumerInterval     = 2 * time.Second
	consumerProcInterval = 300 * time.Millisecond
	consumerWaitInterval = 100 * time.Millisecond
	remainerProcInterval = 1 * time.Second
)

// ProducerConsumer は、ゴルーチンとチャネルを使って 生産者/消費者 処理を実施するサンプルです
func ProducerConsumer() error {
	var (
		// 処理終了を指示するチャネル
		done = make(chan empty)
	)

	// 生産者 生成
	itemCh, termProducer := makeProducer(done)
	// 消費者 生成
	termConsumer := makeConsumer(done, itemCh)
	// 残り物収集班 生成
	termRemainCollect := makeRemainCollector(done, termConsumer, itemCh)

	select {
	case <-time.After(10 * time.Second):
		// 処理終了
		close(done)
	}

	// 生産者, 消費者, 残り物収集班 が終了するのを待つ
	<-termProducer
	<-termConsumer
	<-termRemainCollect

	log.Println("MAIN END")

	return nil
}

// makeProducer は、生産者処理を担当する関数です.
// 内部で、ゴルーチンを起動し以下のスケジューリングで処理を行います。
//
//   - producerInterval毎に一つ生産する
//
// 戻り値として、生産したアイテムを受け取ることができるチャネルと
// 自身の処理が終了したことを知らせるチャネルを返します。
func makeProducer(done done) (itemCh, terminate) {
	ch := make(chan item, 100)
	termCh := make(chan empty)

	go func() {
		defer close(termCh)
		defer close(ch)

		i := 0

		for {
			select {
			case <-done:
				return
			case <-time.After(producerInterval):
				i++
				log.Printf("[生産者] 生成 %d\n", i)
				ch <- item(i)
			}
		}
	}()

	return ch, termCh
}

// makeConsumer は、消費者処理を担当する関数です.
// 内部で、ゴルーチンを起動し以下のスケジューリングで処理を行います。
//
//   - consumerInterval毎に消費処理を実施する
//   - 一度の消費処理で与えられている猶予時間はconsumerProcIntervalとする
//   - 一つのアイテムを消費するのにconsumerWaitIntervalかかる
//
// 戻り値として
// 自身の処理が終了したことを知らせるチャネルを返します。
func makeConsumer(done done, itemCh itemCh) terminate {
	termCh := make(chan empty)

	go func() {
		defer close(termCh)

		for {
			select {
			case <-done:
				return
			case <-time.After(consumerInterval):
				// 消費出来るタイミングが訪れたため活動を開始するが
				// 処理を実施するのは与えられた猶予時間分だけ動くようにする
				timeout := time.After(consumerProcInterval)

			L:
				for {
					select {
					case <-timeout:
						// 消費猶予時間が終わったので今回分のターンは終わり
						break L
					case v, ok := <-itemCh:
						if !ok {
							break L
						}
						log.Printf("[消費者] 消費 %v\n", v)
						time.Sleep(consumerWaitInterval)
					}
				}
			}
		}
	}()

	return termCh
}

// makeRemainCollector は、残り物回収処理を担当する関数です.
// 生産者が生産した分を消費者が消費しきれなかった場合に本処理が残りを回収します。
// 内部で、ゴルーチンを起動し以下のスケジューリングで処理を行います。
//
//   - 管理者(今はmain処理)が処理終了を告げていること
//   - 消費者が処理終了を告げていること
//   - 生産者が生産した残りのアイテムを全て回収する
//   - アイテム一つの回収でremainerProcIntervalかかる
//
// 戻り値として
// 自身の処理が終了したことを知らせるチャネルを返します。
func makeRemainCollector(done done, termConsumer terminate, itemCh itemCh) terminate {
	termCh := make(chan empty)

	go func() {
		defer close(termCh)

		// 自身は残りモノを回収する役割なので
		// 生産終了と消費終了が告げられた後に
		// 動きを開始する
		<-done
		<-termConsumer

		for {
			select {
			case v, ok := <-itemCh:
				if !ok {
					return
				}
				log.Printf("[残り回収班] 回収 %v\n", v)
				time.Sleep(remainerProcInterval)
			}
		}
	}()

	return termCh
}

// Go 1.26
//
// failsafe-go を使って、サーキットブレーカを実現するサンプル。
//
// - 連続失敗 threshold 回で Open
// - Open 中は即座に circuitbreaker.ErrOpen
// - openTimeout 経過後に Half-Open
// - Half-Open で successThreshold 回連続成功したら Closed
// - Half-Open で失敗したら再び Open
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/circuitbreaker"
)

func sendRequest(addr, msg string) error {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.DialTimeout("tcp", addr, 2*time.Second); err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	if err = conn.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		return fmt.Errorf("set deadline: %w", err)
	}

	if _, err = fmt.Fprintf(conn, "%s\n", msg); err != nil {
		return fmt.Errorf("write request: %w", err)
	}

	var (
		reader = bufio.NewReader(conn)
		resp   string
	)
	if resp, err = reader.ReadString('\n'); err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// 改行除去
	resp = strings.TrimSpace(resp)

	// サーバが ERR を返した場合は失敗扱い
	if strings.HasPrefix(resp, "ERR") {
		return fmt.Errorf("server error: %s", resp)
	}

	log.Printf("[client] response: %s", resp)

	return nil
}

func main() {
	log.SetFlags(log.Lmicroseconds)

	var (
		addr             = flag.String("addr", "localhost:9000", "server address")
		requests         = flag.Int("n", 20, "number of requests")
		interval         = flag.Duration("interval", 300*time.Millisecond, "interval between requests")
		failureThreshold = flag.Uint("fail-threshold", 3, "consecutive failures to open circuit")
		successThreshold = flag.Uint("succ-threshold", 2, "consecutive successes to close circuit")
		openTimeout      = flag.Duration("open-timeout", 3*time.Second, "how long circuit stays open")
	)
	flag.Parse()

	// failsafe-go で CircuitBreaker を構築
	//
	// WithFailureThreshold:
	//   Closed 状態で failureThreshold 回連続失敗すると Open へ遷移
	//
	// WithDelay:
	//   Open の維持時間です。経過後に Half-Open へ遷移
	//
	// WithSuccessThreshold:
	//   Half-Open 中に successThreshold 回連続成功すると Closed に戻る
	//
	// OnStateChanged:
	//   状態遷移観測のためのイベント
	var (
		onStateChanged = func(e circuitbreaker.StateChangedEvent) {
			log.Printf("[CB] %s → %s", e.OldState, e.NewState)
		}

		// Circuit Breakerのビルダーを構築
		builder = circuitbreaker.NewBuilder[any]().
			WithFailureThreshold(*failureThreshold).
			WithDelay(*openTimeout).
			WithSuccessThreshold(*successThreshold).
			OnStateChanged(onStateChanged)

		// ビルダーで設定した条件を元に CircuitBreaker を構築
		//
		// 必要なら「どのエラーを failure とみなすか」を絞ることも可能。
		// 今回は sendRequest が返す error をそのまま失敗として扱いたいため
		// 明示的な HandleErrors 設定は省略している。
		breaker = builder.Build()

		// 同じ依存先（同じ TCP サーバ）に対しては breaker を共有するのが推奨パターン。
		// これにより、あるリクエストで Open した状態が他のリクエストにも反映されるため。
		executor = failsafe.With(breaker)
	)
	log.Printf(
		"[client] start addr=%s n=%d fail=%d succ=%d open_timeout=%s",
		*addr,
		*requests,
		*failureThreshold,
		*successThreshold,
		*openTimeout,
	)

	var (
		msg   string
		err   error
		state circuitbreaker.State
	)
	for i := 1; i <= *requests; i++ {
		msg = fmt.Sprintf("req-%03d", i)

		// サーキットブレーカ経由で処理
		//
		// Open 中なら sendRequest は呼ばれず、circuitbreaker.ErrOpen が返る。
		// Half-Open なら 試験的実行 が許可され、その成否で Close/Open が決まる。
		//
		// executor.Get()を使うことも可能。この場合はBuilder生成時に指定した型パラメータが働く。
		err = executor.Run(func() error {
			return sendRequest(*addr, msg)
		})

		// 現在の状態を取得
		state = breaker.State()
		switch {
		case errors.Is(err, circuitbreaker.ErrOpen):
			log.Printf("[client] #%03d %-10s ⚡ CIRCUIT OPEN — リクエストをスキップ", i, state)

		case err != nil:
			log.Printf("[client] #%03d %-10s ✗ FAIL %v", i, state, err)

		default:
			log.Printf("[client] #%03d %-10s ✓ OK", i, state)
		}

		time.Sleep(*interval)
	}

	log.Printf("[client] done final_state=%s", breaker.State())
}

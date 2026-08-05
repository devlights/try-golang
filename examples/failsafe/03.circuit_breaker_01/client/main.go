// client.go
// サーキットブレーカーパターンを標準ライブラリのみで実装した TCP クライアント。
//
// 状態遷移:
//
//	Closed    → 通常リクエスト。連続失敗が threshold を超えると Open へ。
//	Open      → 即座に ErrCircuitOpen を返す。timeout 後 Half-Open へ。
//	Half-Open → 試験リクエストを送信。成功なら Closed、失敗なら Open。
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// ── State ────────────────────────────────────────────

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "Closed"
	case StateOpen:
		return "Open"
	case StateHalfOpen:
		return "Half-Open"
	default:
		return "Unknown"
	}
}

var ErrCircuitOpen = errors.New("circuit is open")

// ── CircuitBreaker ────────────────────────────────────

type CircuitBreaker struct {
	mu sync.Mutex

	// 設定
	failureThreshold int           // この回数連続で失敗したら Open へ
	successThreshold int           // Half-Open 時にこの回数連続で成功したら Closed へ
	openTimeout      time.Duration // Open を維持する時間

	// 状態
	state           State
	consecutiveFail int // 連続失敗数
	consecutiveSucc int // 連続成功数
	openedAt        time.Time
}

func NewCircuitBreaker(failThreshold, succThreshold int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: failThreshold,
		successThreshold: succThreshold,
		openTimeout:      openTimeout,
		state:            StateClosed,
	}
}

// Call はサーキットブレーカーを通じて fn を呼び出す。
// Open 状態では fn を実行せず ErrCircuitOpen を返す。
func (cb *CircuitBreaker) Call(fn func() error) error {
	var (
		s State
	)
	cb.mu.Lock()
	{
		// Open タイムアウトを確認して必要なら Half-Open に遷移
		if cb.state == StateOpen && time.Since(cb.openedAt) >= cb.openTimeout {
			cb.state = StateHalfOpen
			cb.consecutiveFail = 0
			cb.consecutiveSucc = 0
			log.Printf("[CB] Open → Half-Open (試験リクエストを許可)")
		}

		s = cb.state
	}
	cb.mu.Unlock()

	if s == StateOpen {
		return ErrCircuitOpen
	}

	// 処理実行
	var (
		err = fn()
	)
	cb.mu.Lock()
	{
		if err != nil {
			cb.onFailure()
		} else {
			cb.onSuccess()
		}
	}
	cb.mu.Unlock()

	return err
}

func (cb *CircuitBreaker) onSuccess() {
	switch cb.state {
	case StateClosed:
		cb.consecutiveFail = 0
	case StateHalfOpen:
		cb.consecutiveSucc++
		if cb.consecutiveSucc >= cb.successThreshold {
			cb.state = StateClosed
			cb.consecutiveFail = 0
			cb.consecutiveSucc = 0
			log.Printf("[CB] Half-Open → Closed (回路を閉じる)")
		}
	}
}

func (cb *CircuitBreaker) onFailure() {
	switch cb.state {
	case StateClosed:
		cb.consecutiveFail++
		if cb.consecutiveFail >= cb.failureThreshold {
			cb.state = StateOpen
			cb.openedAt = time.Now()
			log.Printf("[CB] Closed → Open (連続失敗 %d 回、%.1fs 遮断)",
				cb.consecutiveFail, cb.openTimeout.Seconds())
		}
	case StateHalfOpen:
		cb.state = StateOpen
		cb.openedAt = time.Now()
		cb.consecutiveSucc = 0
		log.Printf("[CB] Half-Open → Open (試験失敗、再遮断)")
	}
}

func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// ── TCP リクエスト ────────────────────────────────────

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
		return fmt.Errorf("setdeadline: %w", err)
	}

	fmt.Fprintf(conn, "%s\n", msg)

	var (
		resp string
		r    = bufio.NewReader(conn)
	)
	if resp, err = r.ReadString('\n'); err != nil {
		return fmt.Errorf("read: %w", err)
	}

	resp = strings.TrimSpace(resp)
	if strings.HasPrefix(resp, "ERR") {
		return fmt.Errorf("server error: %s", resp)
	}

	log.Printf("[client] response: %s", resp)

	return nil
}

// ── main ──────────────────────────────────────────────

func main() {
	log.SetFlags(log.Lmicroseconds)

	var (
		addr          = flag.String("addr", "localhost:9000", "server address")
		requests      = flag.Int("n", 20, "number of requests")
		interval      = flag.Duration("interval", 300*time.Millisecond, "interval between requests")
		failThreshold = flag.Int("fail-threshold", 3, "consecutive failures to open circuit")
		succThreshold = flag.Int("succ-threshold", 2, "consecutive successes to close circuit")
		openTimeout   = flag.Duration("open-timeout", 3*time.Second, "how long circuit stays open")
	)
	flag.Parse()

	var (
		cb = NewCircuitBreaker(*failThreshold, *succThreshold, *openTimeout)
	)
	log.Printf(
		"[client] start  addr=%s  n=%d  fail=%d  succ=%d  open_timeout=%s",
		*addr,
		*requests,
		*failThreshold,
		*succThreshold,
		*openTimeout)

	for i := 1; i <= *requests; i++ {
		var (
			msg   = fmt.Sprintf("req-%03d", i)
			fn    = func() error { return sendRequest(*addr, msg) }
			err   = cb.Call(fn)
			state = cb.State()
		)
		switch {
		case errors.Is(err, ErrCircuitOpen):
			log.Printf("[client] #%03d %-10s ⚡ CIRCUIT OPEN — リクエストをスキップ", i, state)
		case err != nil:
			log.Printf("[client] #%03d %-10s ✗ FAIL  %v", i, state, err)
		default:
			log.Printf("[client] #%03d %-10s ✓ OK", i, state)
		}

		time.Sleep(*interval)
	}

	log.Printf("[client] done  final_state=%s", cb.State())
}

// RetryPolicy + Timeout を組み合わせて、
//
//   - 「最大3回リトライ」 (RetryPolicy)
//   - 「各試行ごとに1秒タイムアウト」 (TimeoutPolicy)
//   - 「タイムアウトした試行は実処理も止める」(Context)
//
// という挙動を実現するサンプルです。
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
	"github.com/failsafe-go/failsafe-go/timeout"
)

var (
	appLog  *log.Logger
	dbgLog  *log.Logger
	verbose bool
)

func init() {
	appLog = log.New(os.Stderr, "", log.Lmicroseconds)
	dbgLog = log.New(io.Discard, "", log.Lmicroseconds)

	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.Parse()

	if verbose {
		dbgLog.SetOutput(os.Stderr)
	}
}

func main() {
	addr := "localhost:8080"

	// --- RetryPolicy の構築 ---
	//
	// ・WithDelay(time.Second)   : 各リトライ間隔 1 秒
	// ・WithMaxRetries(3)        : 最大 3 回リトライ（= 最大 4 試行）
	//   ※「最大3回リトライ」という日本語を「3回まで再試行」と
	//     解釈しているため、実試行回数としては 1 + 3 = 4 となる。
	//
	retryPolicy := retrypolicy.NewBuilder[net.Conn]().
		WithDelay(time.Second).
		WithMaxRetries(3).
		Build()

	// --- Timeout ポリシーの構築 ---
	//
	// ・各試行ごとに 1 秒のタイムアウトを適用。
	//   Timeout は「その試行に対して子 Context を張り、
	//   時間超過時にキャンセルする」ポリシー。
	//
	timeoutPolicy := timeout.New[net.Conn](time.Second)

	// --- Executor の構築 ---
	//
	// 順序が重要。今回の場合はリトライが外で、タイムアウトが内となる。
	//   failsafe.With(retryPolicy, timeoutPolicy)
	//
	// Timeout が「内側」の場合、各試行ごとに Timeout が適用され
	// Timeout で失敗した試行も、RetryPolicy によって再試行対象になる。
	//
	executor := failsafe.With(retryPolicy, timeoutPolicy)

	// --- 実処理: Dial 処理 ---
	//
	// GetWithExecution を使うことで、Execution を受け取り、
	// Execution.Context() から Timeout が張った子 Context を
	// 取得できる。
	//
	// この Context を DialContext に渡すことで、Timeout 発動時に
	// Dial 自体が中断され、「タイムアウトした試行は実処理も止まる」
	// という挙動になる。
	//
	dialFn := func(exec failsafe.Execution[net.Conn]) (net.Conn, error) {
		conn, err := new(net.Dialer).DialContext(exec.Context(), "tcp4", addr)
		if err != nil {
			err = fmt.Errorf("dial failed: attempt=[%d] retry=[%d] %w", exec.Attempts(), exec.Retries(), err)
			dbgLog.Printf("net.Dialer=> %v\n", err)

			return nil, err
		}

		return conn, nil
	}

	// --- 実行 ---
	//
	// GetWithExecution を用いることで Execution を受け取る関数を渡せる。
	//
	conn, err := executor.GetWithExecution(dialFn)
	if err != nil {
		appLog.Printf("失敗: %[1]v (%[1]T)\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	appLog.Printf("接続に成功しました: %v\n", conn.RemoteAddr())
}

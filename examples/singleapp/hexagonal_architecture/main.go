/*
ヘキサゴナルアーキテクチャのサンプル

adapter -> port <- usecase -> domain

- adapter
  - FixedAmountAdapter (Inbound)
  - ConsoleAdapter     (Outbound)

- port
  - InputPort
  - OutputPort

- usecase
  - YenCalculator

- domain
  - Yen
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ------------------------------------
// ドメインエラー
// ------------------------------------
type DomainError struct {
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func NewDomainError(msg string, cause error) *DomainError {
	return &DomainError{
		Message: msg,
		Cause:   cause,
	}
}

// ------------------------------------
// エンティティ
// ------------------------------------
type Yen struct {
	Amount  int64
	TaxRate float64
}

func (me Yen) WithTax() (int64, error) {
	if me.Amount < 0 {
		return 0, NewDomainError("金額は0以上である必要があります", nil)
	}

	if me.TaxRate <= 0 || me.TaxRate >= 100 {
		return 0, NewDomainError("税率は0より大きく100未満である必要があります", nil)
	}

	v := float64(me.Amount) * me.TaxRate
	return int64(v), nil
}

// ------------------------------------
// ユースケース (ビジネスロジック)
// ------------------------------------
type YenCalculator struct {
	in  InputPort
	out OutputPort
}

func NewYenCalculator(in InputPort, out OutputPort) *YenCalculator {
	return &YenCalculator{in, out}
}

func (me *YenCalculator) Calc() error {
	v, err := me.in.Value()
	if err != nil {
		return fmt.Errorf("入力値の取得に失敗しました: %w", err)
	}

	y := Yen{v, 10.0}
	result, err := y.WithTax()
	if err != nil {
		return fmt.Errorf("税込計算に失敗しました: %w", err)
	}

	if err := me.out.Output(v, result); err != nil {
		return fmt.Errorf("出力処理に失敗しました: %w", err)
	}

	return nil
}

// ------------------------------------
// ポート
// ------------------------------------

// ------------------------------------
// Inbound
type InputPort interface {
	Value() (int64, error)
}

// ------------------------------------
// Outbound
type OutputPort interface {
	Output(before, after int64) error
}

// ------------------------------------
// アダプタ
// ------------------------------------

// ------------------------------------
// プライマリアダプタ
type FixedAmountAdapter struct{}

func (me *FixedAmountAdapter) Value() (int64, error) {
	return int64(100), nil
}

type SquaredAmountAdapter struct {
	Amount int
}

func (me *SquaredAmountAdapter) Value() (int64, error) {
	if me.Amount < 0 {
		return 0, errors.New("入力値は0以上である必要があります")
	}
	v := me.Amount * me.Amount
	return int64(v), nil
}

// 入力エラーをシミュレートするアダプター
type ErrorInputAdapter struct{}

func (me *ErrorInputAdapter) Value() (int64, error) {
	return 0, errors.New("入力ソースからの読み取りに失敗しました")
}

// ------------------------------------
// セカンダリアダプタ
type ConsoleAdapter struct{}

func (me *ConsoleAdapter) Output(before, after int64) error {
	_, err := fmt.Printf("b:%d\ta:%d\n", before, after)
	return err
}

type FileOutAdapter struct {
	Writer io.Writer
}

func (me *FileOutAdapter) Output(before, after int64) error {
	_, err := fmt.Fprintf(me.Writer, "b:%d\ta:%d\n", before, after)
	return err
}

// ------------------------------------
// ユーティリティ
// ------------------------------------
func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)

		// ドメインエラーとその他のエラーで処理を分ける
		var domainErr *DomainError
		if errors.As(err, &domainErr) {
			fmt.Fprintf(os.Stderr, "ドメインエラー: %s\n", domainErr.Message)
		}
	}
}

func main() {
	var (
		in  InputPort      // Inbound Port
		out OutputPort     // Outbound Port
		uc  *YenCalculator // Usecase
	)

	// 1: 基本パターン
	in = &FixedAmountAdapter{}
	out = &ConsoleAdapter{}
	uc = NewYenCalculator(in, out)

	if err := uc.Calc(); err != nil {
		handleError(err)
	} else {
		fmt.Println("計算1が正常に完了しました")
	}

	// 2: 出力アダプターを変更
	out = &FileOutAdapter{Writer: os.Stderr}
	uc = NewYenCalculator(in, out)

	if err := uc.Calc(); err != nil {
		handleError(err)
	} else {
		fmt.Println("計算2が正常に完了しました")
	}

	// 3: 入力アダプターを変更
	in = &SquaredAmountAdapter{Amount: 20}
	uc = NewYenCalculator(in, out)

	if err := uc.Calc(); err != nil {
		handleError(err)
	} else {
		fmt.Println("計算3が正常に完了しました")
	}

	// 4: エラーケース - 入力エラー
	in = &ErrorInputAdapter{}
	uc = NewYenCalculator(in, out)

	if err := uc.Calc(); err != nil {
		handleError(err)
	} else {
		fmt.Println("計算4が正常に完了しました")
	}

	// 5: エラーケース - ドメインエラー（負の入力値）
	in = &SquaredAmountAdapter{Amount: -10}
	uc = NewYenCalculator(in, out)

	if err := uc.Calc(); err != nil {
		handleError(err)
	} else {
		fmt.Println("計算5が正常に完了しました")
	}
}

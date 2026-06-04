package main

import (
	"context"
	"errors"
	"log"
	"time"
)

const (
	MainTimeout = 100 * time.Millisecond
	ProcTimeout = 500 * time.Microsecond
)

var (
	ErrMainTooSlow = errors.New("[MAIN] TOO SLOW")
	ErrProcTooSlow = errors.New("[PROC] TOO SLOW")
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, MainTimeout, ErrMainTooSlow)
		procCtx          = run(mainCtx)
		err              error
	)
	defer mainCxl()

	select {
	case <-mainCtx.Done():
		err = context.Cause(mainCtx)
	case <-procCtx.Done():
		if err = context.Cause(procCtx); errors.Is(err, context.Canceled) {
			err = nil
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func run(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancelCause(pCtx)
	)

	go func() {
		cxl(proc(ctx))
	}()
	go func() {
		<-time.After(ProcTimeout)
		cxl(ErrProcTooSlow)
	}()

	return ctx
}

func proc(_ context.Context) error {
	const (
		bitSize = 32
	)

	//
	// ビットの設定
	//
	var (
		bv  = NewBitVector(bitSize)
		err error
	)

	for _, i := range []int{1, 3, 23} {
		if err = bv.Set(i, true); err != nil {
			return err
		}
	}
	log.Printf("ORIG : %s", bv)

	//
	// ビットの取得
	//
	var (
		bits []bool
		bit  bool
	)

	for _, i := range []int{1, 2, 3, 4} {
		if bit, err = bv.Get(i); err != nil {
			return err
		}

		bits = append(bits, bit)
	}
	log.Printf("BITS : %v", bits)

	//
	// 別のビットベクタとの演算
	//
	var (
		bv2   = NewBitVector(bitSize)
		bvAnd *BitVector
		bvOr  *BitVector
		bvXor *BitVector
	)

	for _, i := range []int{2, 23, 24} {
		if err = bv2.Set(i, true); err != nil {
			return err
		}
	}
	log.Printf("OTHER: %s", bv2)

	// AND
	if bvAnd, err = bv.And(bv2); err != nil {
		return err
	}
	log.Printf("AND  : %s", bvAnd)

	// OR
	if bvOr, err = bv.Or(bv2); err != nil {
		return err
	}
	log.Printf("OR   : %s", bvOr)

	// XOR
	if bvXor, err = bv.Xor(bv2); err != nil {
		return err
	}
	log.Printf("XOR  : %s", bvXor)

	return nil
}

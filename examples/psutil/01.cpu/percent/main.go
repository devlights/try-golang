package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"math"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type (
	Args struct {
		perCpu  bool
		spinCpu bool
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.perCpu, "percpu", false, "per-cpu")
	flag.BoolVar(&args.spinCpu, "spincpu", false, "spin-cpu")
}

func main() {
	log.SetFlags(log.Ltime)
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	type (
		cpuval struct {
			val []int
			err error
		}
	)
	var (
		interval = 1 * time.Second
		percpu   = args.perCpu
		cpuCh    = make(chan *cpuval)
		ctx, cxl = context.WithTimeout(context.Background(), 6*time.Second)
	)
	defer cxl()
	defer close(cpuCh)

	// 必要であれば無駄にCPUを回す
	if args.spinCpu {
		go spinCpu(ctx)
	}

	// CPU使用率を取得
	go func(ctx context.Context, ch chan<- *cpuval) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				v, err := cpu.PercentWithContext(ctx, interval, percpu)

				vals := make([]int, len(v))
				for i, f := range v {
					vals[i] = int(f)
				}

				cpuCh <- &cpuval{vals, err}
			}
		}
	}(ctx, cpuCh)

	// 出力
	for v := range cpuCh {
		if v.err != nil {
			if errors.Is(v.err, context.DeadlineExceeded) {
				return nil
			}
			return v.err
		}

		log.Printf("%v\n", v.val)
	}

	return nil
}

func spinCpu(ctx context.Context) {
	for range runtime.NumCPU() {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					for i := range 100000 {
						_ = math.Sqrt(float64(i * i))
					}
					time.Sleep(1 * time.Microsecond)
				}
			}
		}(ctx)
	}
}

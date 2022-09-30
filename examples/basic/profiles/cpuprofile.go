package profiles

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime/pprof"
	"time"

	"github.com/devlights/gomy/errs"
	"github.com/devlights/gomy/output"
)

// CpuProfile は、pprof を使ってCPUプロファイルを取得するサンプルです。
func CpuProfile() error {
	const (
		ShellPath   = "/bin/bash"
		ShellOpt    = "-c"
		ProfilePath = "examples/basic/profiles/cpu.pprof"
	)

	// --------------------------------------------------------
	// Make contexts
	// --------------------------------------------------------

	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 10*time.Second)
	)
	defer cxl()

	// --------------------------------------------------------
	// Collect CPU profiles
	// --------------------------------------------------------

	var (
		f = errs.Panic(os.Create(ProfilePath))
	)

	pprof.StartCPUProfile(f)
	func() {
		defer f.Close()
		defer pprof.StopCPUProfile()
		<-run(ctx).Done()
	}()

	output.StdoutHr()

	// --------------------------------------------------------
	// Launch pprof-tool
	// --------------------------------------------------------

	var (
		cmd = exec.CommandContext(ctx, ShellPath, ShellOpt, fmt.Sprintf("go tool pprof -top %s", ProfilePath))
		buf = errs.Panic(cmd.CombinedOutput())
	)

	output.Stdoutl("", string(buf))

	return nil
}

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
//
// # References
//
//   - https://pkg.go.dev/runtime/pprof@go1.19.1
//   - https://github.com/pkg/profile/blob/master/profile.go
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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: profiles_cpu

	   [Name] "profiles_cpu"
	   https://github.com/devlights/try-csharp [170039bytes]
	   https://github.com/devlights/try-golang [170745bytes]
	   https://devlights.hatenablog.com/ [121967bytes]
	   https://github.com/devlights/gomy [168467bytes]
	   https://github.com/devlights/try-python [182642bytes]
	   https://github.com/devlights/goxcel [170197bytes]
	   --------------------------------------------------
	   File: try-golang
	   Type: cpu
	   Time: Jan 30, 2024 at 2:32am (UTC)
	   Duration: 889.51ms, Total samples = 10ms ( 1.12%)
	   Showing nodes accounting for 10ms, 100% of 10ms total
	         flat  flat%   sum%        cum   cum%
	         10ms   100%   100%       10ms   100%  runtime/internal/syscall.Syscall6
	            0     0%   100%       10ms   100%  crypto/tls.(*Conn).HandshakeContext (inline)
	            0     0%   100%       10ms   100%  crypto/tls.(*Conn).clientHandshake
	            0     0%   100%       10ms   100%  crypto/tls.(*Conn).flush
	            0     0%   100%       10ms   100%  crypto/tls.(*Conn).handshakeContext
	            0     0%   100%       10ms   100%  crypto/tls.(*clientHandshakeStateTLS13).handshake
	            0     0%   100%       10ms   100%  internal/poll.(*FD).Write
	            0     0%   100%       10ms   100%  internal/poll.ignoringEINTRIO (inline)
	            0     0%   100%       10ms   100%  net.(*conn).Write
	            0     0%   100%       10ms   100%  net.(*netFD).Write
	            0     0%   100%       10ms   100%  net/http.(*persistConn).addTLS.func2
	            0     0%   100%       10ms   100%  syscall.RawSyscall6
	            0     0%   100%       10ms   100%  syscall.Syscall
	            0     0%   100%       10ms   100%  syscall.Write (inline)
	            0     0%   100%       10ms   100%  syscall.write



	   [Elapsed] 900.186041ms
	*/

}

test:
	go test .

test_verbose:
	go test -v .

coverage:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out

bench:
	go test . -bench .

bench_only:
	go test . -bench . -run ^$$

bench_mem:
	go test . -bench . -benchmem -run ^$$

profile:
	go test . -bench . -benchmem -run ^$$ -memprofile mem.pprof -cpuprofile cpu.pprof

pprof_cpu:
	go tool pprof cpu.pprof

pprof_mem:
	go tool pprof mem.pprof
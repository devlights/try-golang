#!/usr/bin/env bash

go install github.com/go-task/task/v3/cmd/task@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/perf/cmd/benchstat@latest

sudo apt update && sudo sudo apt install -y universal-ctags tree nkf wamerican miller tcpdump
wget -O /tmp/hyperfine.deb https://github.com/sharkdp/hyperfine/releases/download/v1.19.0/hyperfine_1.19.0_amd64.deb
sudo dpkg -i /tmp/hyperfine.deb

go mod download
task build
#!/usr/bin/env bash

go install github.com/go-task/task/v3/cmd/task@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-delve/delve/cmd/dlv@latest

sudo apt update && sudo apt install -y universal-ctags

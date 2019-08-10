# try-golang
This is my TUTORIAL project for golang

[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge)](https://www.codefactor.io/repository/github/devlights/try-golang)

## Environment

```sh
$ sw_vers 
ProductName:	Mac OS X
ProductVersion:	10.14.4
BuildVersion:	18E226
```

## GO version

```sh
$ go version
go version go1.12 darwin/amd64
```

## IDE
- JetBrains GoLand

## Run

```sh
$ go run cmd/trygolang/main.go
```

## Test

```sh
$ go test github.com/devlights/try-golang/...
```

## Install (executable module only)

```sh
$ go get -u github.com/devlights/try-golang/cmd/trygolang
$ cd $(go env GOPATH)/bin
$ trygolang
```
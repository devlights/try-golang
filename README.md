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
go version go1.13 darwin/amd64
```

## IDE
- JetBrains GoLand

## Run

```sh
$ go run github.com/devlights/cmd/trygolang/*.go
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
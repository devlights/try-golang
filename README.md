# try-golang
This is my TUTORIAL project for golang

![try-golang - Go Version](https://img.shields.io/badge/go-1.13-blue.svg)
[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge)](https://www.codefactor.io/repository/github/devlights/try-golang)

## GO version

```shell script
$ go version
go version go1.13.4 darwin/amd64
```

```shell script
$ go version
go version go1.13.4 windows/amd64
```

## IDE

- JetBrains GoLand

## Run

```shell script
$ cd cmd/trygolang
$ go run .
```

or 

```shell script
$ make run
```

## Test

```shell script
$ go test -v ./...
```

or

```shell script
$ make test
```

## Install (executable module only)

```sh
$ go get -u -v github.com/devlights/try-golang/cmd/trygolang
$ cd $(go env GOPATH)/bin
$ trygolang
```

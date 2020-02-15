
# try-golang
This is my TUTORIAL project for golang

![try-golang - Go Version](https://img.shields.io/badge/go-1.13-blue.svg)
[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge)](https://www.codefactor.io/repository/github/devlights/try-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/devlights/try-golang)](https://goreportcard.com/report/github.com/devlights/try-golang)

## GO version

```shell script
$ go version
go version go1.13.6 linux/amd64
```

```shell script
$ go version
go version go1.13.6 darwin/amd64
```

```shell script
$ go version
go version go1.13.6 windows/amd64
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

or

```shell script
$ make run EXAMPLE=example_name
```

## Test

```shell script
$ go test -v ./...
```

or

```shell script
$ make test
```

## Install

```shell script
$ go env -w GO111MODULE=off
$ go get -u github.com/devlights/try-golang/cmd/trygolang
$ go env -u GO111MODULE
$ cd $(go env GOPATH)/bin
$ trygolang
```

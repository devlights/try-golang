# try-golang

This is my TUTORIAL project for golang

![try-golang - Go Version](https://img.shields.io/badge/go-1.16-blue.svg)
[![CodeFactor](https://www.codefactor.io/repository/github/devlights/try-golang/badge/master)](https://www.codefactor.io/repository/github/devlights/try-golang/overview/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/devlights/try-golang)](https://goreportcard.com/report/github.com/devlights/try-golang)
![Go](https://github.com/devlights/try-golang/workflows/Go/badge.svg?branch=master)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/devlights/try-golang) 
[![PkgGoDev](https://pkg.go.dev/badge/github.com/devlights/try-golang)](https://pkg.go.dev/github.com/devlights/try-golang)

## GO version

```shell script
$ go version
go version go1.16 linux/arm64
```

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

or

```shell script
$ make docker
```

### Run Docker on Gitpod

type following command in first terminal:

```shell script
$ sudo docker-up
```

Launch new terminal and type following command:

```shell
$ make docker
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
$ cd cmd/trygolang
$ go install
```

or 

```shell script
$ make install
```

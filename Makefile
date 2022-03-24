EXAMPLE=""

all: clean build test

prepare:
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/go-task/task/v3/cmd/task@latest

build: 
	go build

build-static: 
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"'

test: 
	go test -coverprofile /tmp/try-golang-cover $(shell go list ./... | grep -v /examples/ | grep -v /cmd/)

clean: 
	go clean

install: 
	go install

vet:
	go vet ./...
	staticcheck ./...

run:
	go run main.go -onetime -example ${EXAMPLE}

generate: 
	go generate -x ./...

docker-build:
	DOCKER_BUILDKIT=1 docker image build -t try-golang -f Dockerfile ${PWD}

docker-run: docker-build
	docker container run -it --rm --name try-golang try-golang

docker-sh: docker-build
	docker container run -it --rm --name try-golang try-golang bash

docker: docker-run


default: test-normal test-short

prepare:
	@go clean -testcache

test-normal: prepare
	@echo '-------------------------------------------------------'
	go test -v ./lib

test-short: prepare
	@echo '-------------------------------------------------------'
	go test -v -short ./lib

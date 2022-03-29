default: gen run-all

gen:
	tr -dc "A-Za-z 0-9" < /dev/urandom | fold -w1024 | head -n 1000000 > /tmp/try-golang-big.txt
	wc -l /tmp/try-golang-big.txt

run-all: run-nobuffering run-buffering1 run-buffering2

run-nobuffering:
	go run nobuffering/main.go > /dev/null

run-buffering1:
	go run buffering1/main.go > /dev/null

run-buffering2:
	go run buffering2/main.go > /dev/null

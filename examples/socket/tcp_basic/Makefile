all: build start-server run-client stop-server

build:
	@(cd client; go build)
	@(cd server; go build)

start-server:
	@(cd server; ./server & echo $$! > server.PID)

run-client:
	@(cd client; ./client)

stop-server:
	@kill $$(cat server/server.PID)

clean:
	@(cd client; go clean)
	@(cd server; go clean)
	@find . -type f -name "*.PID" -exec rm -rf {} \;
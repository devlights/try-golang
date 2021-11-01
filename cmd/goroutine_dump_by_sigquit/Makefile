build:
	go build -o App main.go

run-foreground: build
	./App
# このあとに Ctrl+¥ (バックスラッシュ) を押す

run-background: build
	./App &
	sleep 1
	pkill --signal SIGQUIT -f ./App

run-docker: build
	docker image build -t app -f Dockerfile ${PWD}
	docker container run -dit --rm --name app001 app
	docker container logs -f app001 > container.log &
	sleep 1
	docker container kill --signal=SIGQUIT app001
#	kill $(shell pgrep -f -o "docker container logs")
	cat container.log

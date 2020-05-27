.PHONY: lint, test, fmtall, all, build, run, re

lint:
	golangci-lint run

test:
	go test -count=1 ./... -cover

fmtall:
	go fmt ./...

all:
	make fmtall
	make lint
	make test

build:
	go build -o restarter

run:
	make build
	nohup ./restarter &

re:
	pid := $(pidof ./restarter)
	kill -9 $(pid)
	rm -f nohup.out
	make run
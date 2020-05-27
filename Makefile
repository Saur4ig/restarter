.PHONY: lint, test, fmtall, all, build, run

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
	./restarter

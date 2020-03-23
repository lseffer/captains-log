.PHONY: build test

build:
	mkdir -p bin
	go build -o bin/

test:
	go test -cover -v -race ./...

.PHONY: run build test

build:
	rm -rf dist
	mkdir -p dist
	cp -r views static sql dist/
	go build -o dist/

run: build
	./dist/captains-log

test:
	go test -cover -v -race ./...

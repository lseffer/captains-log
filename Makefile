.PHONY: run build test

build_setup:
	@rm -rf dist
	@mkdir -p dist
	@cp -r views static sql dist/

build: build_setup
	@go build -o dist/

build_linux_arm: build_setup
	 @docker run --rm -v "$$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=arm golang:1.14 go install github.com/mattn/go-sqlite3 && go build -o dist/

run: build
	@./dist/captains-log -a

test:
	@go test -tags testsuite -cover -coverprofile=coverage.out -race  ./...
	@go tool cover -func=coverage.out

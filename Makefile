.PHONY: run build test

build:
	@rm -rf dist
	@mkdir -p dist
	@cp -r views static sql dist/
	@go build -o dist/

run: build
	@./dist/captains-log -a

test:
	@go test -tags testsuite -cover -coverprofile=coverage.out -race  ./...
	@go tool cover -func=coverage.out

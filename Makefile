build: test ## Make a macOS executable binary
	go build -o bin/mac/goids .
.PHONY: build

build-win: test ## Make a Windows executable binary
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/win/goids.exe .
.PHONY: build-win

test: deps ## go test
	go test ./...
.PHONY: test

deps: fmt ## go mod tidy
	go mod tidy
.PHONY: deps

fmt: ## go fmt
	go fmt
.PHONY: fmt

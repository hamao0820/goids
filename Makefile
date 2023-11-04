build: test build-win build-intel build-m1 lipo clean
.PHONY: build

clean:
	rm bin/mac/goids-intel bin/mac/goids-m1
.PHONY: clean

lipo:
	lipo -create -output bin/mac/goids bin/mac/goids-intel bin/mac/goids-m1
.PHONY: lipo

build-intel:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o bin/mac/goids-intel -ldflags="-s -w" -trimpath .
.PHONY: build-intel

build-m1:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o bin/mac/goids-m1 -ldflags="-s -w" -trimpath .
.PHONY: build-m1

build-win:
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/win/goids.exe -ldflags="-s -w" -trimpath .
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

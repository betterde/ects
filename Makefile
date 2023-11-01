BINARY_NAME ?= ects

export CGO_ENABLED = 0
export GOARCH = amd64

all: test build release

fmt:
	go fmt ./...

test:
	CGO_ENABLED=1 go test -race ./...

build:
	cd web && yarn build
	GOOS=darwin go build -ldflags "-s -w" -o "bin/$(BINARY_NAME)_darwin" main.go
	GOOS=linux go build -ldflags "-s -w" -o "bin/$(BINARY_NAME)_linux" main.go
	GOOS=windows go build -ldflags "-s -w" -o "bin/$(BINARY_NAME)_windows" main.go

install:
	cd web && yarn install

clean:
	go clean -testcache

release:
	goreleaser --rm-dist


BINARY_NAME ?= ects

export CGO_ENABLED = 0
export GOARCH = amd64

build:
	GOOS=darwin go build -o "bin/$(BINARY_NAME)_darwin" main.go
	GOOS=linux go build -o "bin/$(BINARY_NAME)_linux" main.go
	GOOS=windows go build -o "bin/$(BINARY_NAME)_windows" main.go

export GITHUB_TOKEN = `dc959301220984be9588ae38e911121397dd30ad`

release:

# Install golang - brew install go
# Install golint - go get -u golang.org/x/lint/golint

all: build
.PHONY: all

GOPATH = $(shell go env GOPATH)

build:
	$(GOPATH)bin/golint apps udpserver webserver
	GOOS=darwin GOARCH=amd64 go build -o rtm apps/rtm.go

clean:
	go clean
	rm rtm

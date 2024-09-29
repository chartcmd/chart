GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chart
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	$(GOBUILD) -o out/bin/$(BINARY_NAME) -v ./cmd

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f out/bin/$(BINARY_NAME)
	rm -f out/bin/$(BINARY_UNIX)

run:
	$(GOBUILD) -o out/bin/$(BINARY_NAME) -v ./cmd
	./out/bin/$(BINARY_NAME)

deps:
	$(GOGET) ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o out/bin/$(BINARY_UNIX) -v ./cmd

.PHONY: all build test clean run deps build-linux

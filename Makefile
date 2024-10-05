GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chart
BINARY_UNIX=$(BINARY_NAME)

init:
	sudo mkdir ~/.config
	sudo mkdir ~/.config/$(BINARY_NAME)
	sudo touch ~/.config/$(BINARY_NAME)/config.json

build:
	$(GOBUILD) -o out/bin/$(BINARY_NAME) -v ./cmd

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o out/bin/$(BINARY_UNIX) -v ./cmd

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f out/bin/$(BINARY_NAME)
	rm -f out/bin/$(BINARY_UNIX)

reset:
	sudo rm -rf ~/.$(BINARY_NAME)

run:
	$(GOBUILD) -o out/bin/$(BINARY_NAME) -v ./cmd
	./out/bin/$(BINARY_NAME)

deps:
	$(GOCMD) mod tidy

vet:
	$(GOCMD) vet ./...

fmt:
	$(GOCMD) fmt ./...

prod-mac:
	$(GOBUILD) $(LDFLAGS) -o out/bin/$(BINARY_NAME) -v ./cmd
	strip out/bin/$(BINARY_NAME)
	sudo mv out/bin/$(BINARY_NAME) /usr/local/bin/
	sudo chmod +x /usr/local/bin/$(BINARY_NAME)

prod-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o out/bin/$(BINARY_UNIX) -v ./cmd
	strip out/bin/$(BINARY_UNIX)
	sudo mv out/bin/$(BINARY_UNIX) /usr/local/bin/$(BINARY_NAME)
	sudo chmod +x /usr/local/bin/$(BINARY_NAME)

.PHONY: build test clean run deps build-linux prod-mac prod-linux

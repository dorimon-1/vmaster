# Makefile for building Go applications with main.go in ./cmd

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=vmaster
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_DIR=./cmd

all: test build
build: 
	cd $(MAIN_DIR) && $(GOBUILD) -o ../$(BINARY_NAME) 
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	cd $(MAIN_DIR) && $(GOBUILD) -o ../$(BINARY_NAME) -v .
	cd .. && ./$(BINARY_NAME)
deps:
	$(GOGET) -v -d ./...
install: build
	cd ..  
	/bin/bash -c 'sudo cp $(BINARY_NAME) /usr/local/bin/'
	@echo "Installed $(BINARY_NAME) to /usr/local/bin/"

uninstall:
	cd .. && sudo rm /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstalled $(BINARY_NAME) from /usr/local/bin"
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd $(MAIN_DIR) && $(GOBUILD) -o ../$(BINARY_UNIX) -v

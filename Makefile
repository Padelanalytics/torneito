# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
BINARY_NAME=torneito

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME)

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe

test:
	$(GOTEST) -v ./...

test-coverage:
	$(GOTEST) -v -coverprofile coverage.out  ./...

install:
	$(GOINSTALL) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)



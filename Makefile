# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

P_BUILD=dist/
P_CMD=cmd/

all: test build

build:
	$(GOBUILD) -o $(P_BUILD)passemble $(P_CMD)passemble/main.go
	$(GOBUILD) -o $(P_BUILD)pcc $(P_CMD)pcc/main.go
	$(GOBUILD) -o $(P_BUILD)pemu $(P_CMD)pemu/main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(P_BUILD)passemble
	rm -f $(P_BUILD)pcc
	rm -f $(P_BUILD)pemu

run:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v
	./dist/$(BINARY_NAME)

deps:
	echo "No dependencies..."

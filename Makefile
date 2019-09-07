GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=auth-service

all: test build-docs build

test:
	$(GOTEST) -v ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) ./cmd/auth/main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/auth/main.go
	./$(BINARY_NAME)

deps:
	$(GOGET) get github.com/go-swagger/go-swagger/cmd/swagger
	$(GOGET) -u ./...

docs: build-docs serve-docs

build-docs:
	swagger generate spec -o api/openapi-spec/swagger.json

serve-docs:
	swagger serve api/openapi-spec/swagger.json


GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=auth-service

all: test build-docs build

test:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic -v ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) ./cmd/auth/main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/auth/main.go
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/go-swagger/go-swagger/cmd/swagger
	$(GOGET) -u ./...

docs:
	swag init -g cmd/auth/main.go -o api/openapi-spec/

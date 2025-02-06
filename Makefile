PROJECT_NAME = relay	
GO_FILES = $(wildcard *.go) 			
BINARY = bin/$(PROJECT_NAME)

.PHONY: all build run clean tidy buildw

all: clean tidy build run

tidy:
	@echo mod tidy
	@go mod tidy

build:
	@echo building
	@go build -o $(BINARY) $(GO_FILES)

run:
	@echo running
	@$(BINARY)

clean:
	@echo cleaning files
	@del /f /q bin\*

test:
	go test ./... -v

buildw:
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/$(PROJECT_NAME) $(GO_FILES)
	set GOOS=
	set GOARCH=

recreate:
	@$(BINARY) -recreate

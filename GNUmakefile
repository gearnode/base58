BIN_DIR = $(CURDIR)/bin

all: test

vet:
	go vet ./...

test: FORCE
	go test -cover -race ./...

FORCE:

.PHONY: all vet test

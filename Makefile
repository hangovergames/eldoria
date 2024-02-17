.PHONY: build clean tidy

ELDORIA_CLIENT_SOURCES := $(shell find ./cmd/eldoria-client ./internal/client ./internal/common -type f -iname '*.go' ! -iname '*_test.go')

ELDORIA_SERVER_SOURCES := $(shell find ./cmd/eldoria-server ./internal/server ./internal/common -type f -iname '*.go' ! -iname '*_test.go')

all: build

build: eldoria-server eldoria-client

tidy:
	go mod tidy

eldoria-server: $(ELDORIA_SERVER_SOURCES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o eldoria-server ./cmd/eldoria-server
	chmod 700 ./eldoria-server

eldoria-client: $(ELDORIA_CLIENT_SOURCES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o eldoria-client ./cmd/eldoria-client
	chmod 700 ./eldoria-client

test:
	go test -v ./...

clean:
	rm -f eldoria-server eldoria-client

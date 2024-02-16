.PHONY: build clean tidy

ELDORIA_SOURCES := \
    ./cmd/eldoria-server/main.go

all: build

build: eldoria-server

tidy:
	go mod tidy

eldoria-server: $(ELDORIA_SOURCES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o eldoria-server ./cmd/eldoria-server
	chmod 700 ./eldoria-server

test:
	go test -v ./...

clean:
	rm -f eldoria-server

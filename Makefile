VERSION=$(shell git describe --tags 2>/dev/null || git describe --always 2>/dev/null)

all: build

build:
	go build -ldflags "-X main.version=${VERSION}" ./cmd/refract-api-server

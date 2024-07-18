BIN_NAME=go-portal
VERSION=$(shell git describe --tags --always --long)
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')



LINUX=$(BIN_NAME)_linux_amd64.out


$(LINUX): .
	go build -v -o $(LINUX) -tags linux -ldflags="-s -w -X main.Version=$(VERSION) -X main.Time=$(BUILD_DATE)" .

clean:
	rm -f $(LINUX)

build:
	$(LINUX)

test:
	go test -v ./...

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build test clean

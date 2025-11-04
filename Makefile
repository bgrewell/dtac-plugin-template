APP={{plugin}}

.PHONY: all build test run lint tidy clean
all: build

build:
	GO111MODULE=on go build -o bin/$(APP) ./cmd/$(APP)

test:
	go test ./...

run: build
	./bin/$(APP)

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || echo "golangci-lint not installed"

tidy:
	go mod tidy

clean:
	rm -rf bin


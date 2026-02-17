MAIN_FILE = main.go

.PHONY: run dev test fmt

run:
	go run $(MAIN_FILE)

dev: run

test:
	go test ./...

fmt:
	gofmt -w main.go $(shell find api internal -name '*.go')

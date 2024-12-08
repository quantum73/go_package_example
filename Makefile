.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run ./...
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

run: vet
	go run ./...
.PHONY:run

tests:
	go test ./...
.PHONY:tests
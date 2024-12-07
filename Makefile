.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

run: vet
	go run ./...
.PHONY:run

tests:
	go test ./...
.PHONY:tests
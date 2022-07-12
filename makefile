all: run

run:
	go run ./cmd/main.go

test:
	go test ./...

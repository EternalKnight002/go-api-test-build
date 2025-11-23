run:
	go run ./cmd/server

build:
	go build -o server ./cmd/server

fmt:
	go fmt ./...

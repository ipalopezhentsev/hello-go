.DEFAULT_GOAL=run

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
run: vet
	go run .
build: vet
	go build
clean:
	go clean
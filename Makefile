.DEFAULT_GOAL=run

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
run: vet
	go run .
build: vet
	go build -o dist/hello
test: build
	go test
clean:
	go clean
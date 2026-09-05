test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

example:
	go run ./examples/guard

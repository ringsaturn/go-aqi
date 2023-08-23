test:
	golangci-lint run ./...
	go test -race ./...

bench:
	go test -bench=. ./...

stringer:
	# go install golang.org/x/tools/cmd/stringer
	stringer -type=Pollutant

tidy:
	rm go.sum
	go mod tidy

fmt:
	go fmt ./...

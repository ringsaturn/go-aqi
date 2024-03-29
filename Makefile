test:
	golangci-lint run ./...
	go test -race ./...

bench:
	go test -bench=. ./...

stringer:
	# go install golang.org/x/tools/cmd/stringer
	stringer -type=Pollutant

fmt:
	go fmt ./...

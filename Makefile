test:
	golangci-lint run api/... cmd/... internal/...
	go test -race ./...

bench:
	go test -bench=. ./...

tidy:
	rm go.sum
	go mod tidy

.PHONY:pb
pb:
	protoc -I=./ --go_out=./ ./*.proto

fmt:
	find ./ -iname *.proto | xargs clang-format -i --style=Google
	go fmt ./...

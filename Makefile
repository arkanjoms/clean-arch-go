
prepare:
	go get -u github.com/ory/go-acc
	go mod download
	go mod tidy

run:
	go run main.go

fmt:
	go fmt ./...

test: mock
	go-acc ./...

cover:
	go tool cover -html coverage.txt

mock:
	go generate -v ./...

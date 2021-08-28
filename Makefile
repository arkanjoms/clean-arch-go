
prepare:
	go mod download
	go mod tidy

run:
	go run main.go

fmt:
	go fmt ./...

test: mock
	go test ./... -coverprofile=coverage.out

cover:
	go tool cover -html coverage.out

mock:
	go generate -v ./...

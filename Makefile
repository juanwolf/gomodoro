BINARY_NAME="gomodoro"

test:
	go test -cover -race ./...

build:
	go build -o $(BINARY_NAME) main.go
run:
	go run main.go

install: build
	mv gomodoro /usr/lib/bin/

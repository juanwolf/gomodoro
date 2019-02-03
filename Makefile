BINARY_NAME="tomato"

test:
	go test ./...

build:
	go build -o $(BINARY_NAME) main.go
run:
	go run main.go

install: build
	mv tomato /usr/lib/bin/

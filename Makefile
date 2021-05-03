go:
	clear
	go build . && ./webstatus

run:
	go run .

build:
	go build .

test:
	go test -v ./...

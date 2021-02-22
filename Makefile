benchmark:
	go test -bench

test:
	go test

all:
	go build -ldflags="-s -w"

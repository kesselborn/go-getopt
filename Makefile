build:
	gofmt -s=true -w *.go
	go version > TESTED_GO_RELEASE
	go build -x

test:
	go test -v

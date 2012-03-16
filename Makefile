build: fmt
	go version > TESTED_GO_RELEASE
	go build -x

fmt:
	gofmt -s=true -w *.go

test:
	go test -v

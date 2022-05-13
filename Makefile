stage=dev
m_gopath=$(shell go env GOPATH 2>/dev/null || echo "/var/go")

setup:
	mkdir "bin"

build-and-deploy:
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_books_handler *.go
	serverless deploy -s ${stage}
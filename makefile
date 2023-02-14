BINARY_NAME=netlab

build:
	go build -o ${BINARY_NAME} main.go

install:
	go install

build_all:
	GOARCH=amd64 GOOS=darwin go build -o dist/${BINARY_NAME}-macos main.go
	GOARCH=amd64 GOOS=linux go build -o dist/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o dist/${BINARY_NAME}-windows main.go

clean:
	go clean
	rm ${BINARY_NAME}
	rm -rf dist
BINARY_NAME=avalara
API_KEY=6a7f8ba5-7fd5-4911-b90b-fc6fd92b7ef5

build:
	# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME} main.go
	# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} main.go

run: build
	API_KEY=${API_KEY} ./${BINARY_NAME} -word=${word}

test: build
	# API_KEY=${API_KEY} GOOS=linux GOARCH=amd64 go test -v *.go
	API_KEY=${API_KEY} GOOS=windows GOARCH=amd64 go test -v *.go
	# API_KEY=${API_KEY} GOOS=darwin GOARCH=amd64 go test -v *.go

clean:
	go clean
	rm ${BINARY_NAME}

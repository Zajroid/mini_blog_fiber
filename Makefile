BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

test:
	go test -v main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go get github.com/gofiber/fiber/v2


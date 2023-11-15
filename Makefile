BINARY=engine
dev: generate
	go run github.com/cosmtrek/air

run: generate
	go run .

build:
	go build -o ${BINARY} .

generate:
	go generate ./...

.PHONY: dev run build generate

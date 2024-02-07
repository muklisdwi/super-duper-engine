BINARY=engine
dev: generate
	go run github.com/cosmtrek/air

run: generate
	go run .

build:
	go build -o ${BINARY} .

generate:
	go generate ./...

new_migration:
	@echo "example: migrate create -ext sql -dir db/migration -seq name=foobarbaz"
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: dev run build generate new_migration

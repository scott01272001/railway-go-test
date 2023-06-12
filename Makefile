build:
	@go build -o bin/go-gin-gorm

run: build
	@clear
	@./bin/go-gin-gorm

test:
	@go test -v ./...
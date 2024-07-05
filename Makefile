run: build
	@./bin/Send-Email

build:
	@go build -o bin/Send-Email

test:
	@go test -v ./...
run:
	@go run .
build:
	@go build -o .
clean:
	@go rm bin/go-boilerplate
dev:
	@air .
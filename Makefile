build:
	@echo "Buildding application"
	@go build

test:
	@echo "Running test"
	@go test ./...

lint:
	@echo "Running Linter"
	@staticcheck ./...

fmt:
	@echo "Formatting"
	@go fmt ./...

tidy:
	@echo "Running 'go mod tidy'"
	@go mod tidy

release:
	@echo "Release application"
	@echo "test running"
	@go test ./...
	@echo "buildding application"
	@go build
	@echo "build complete"

info:
	@echo "Name : mcpy"
	@echo "Repo : https://github.com/aqyuki/mcpy"
	@echo "Quick start : go run main.go"
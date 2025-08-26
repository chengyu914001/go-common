SERVICE_NAME=test
PREFIX=SERVICE_NAME=$(SERVICE_NAME)

lint:
	@go fmt ./...
.PHONY: lint

local-run:
	@ENV_MODE=local $(PREFIX) go run cmd/main.go
.PHONY: local-run


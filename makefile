LOAD_ENV=env $(cat .env | xargs)

lint:
	@go fmt ./...
.PHONY: lint

test:
	@$(LOAD_ENV) go test -count=1 -v ./...
.PHONY: test

local-run:
	@$(LOAD_ENV) go run cmd/server/main.go
.PHONY: local-run

include .env

PG_DSN := "postgresql://$(STORE_USER):$(STORE_PASSWORD)@$(STORE_HOST):$(STORE_PORT)/$(STORE_DATABASE)?sslmode=disable"

migrate-up:
	go tool goose -dir migrations/ postgres $(PG_DSN) up

migrate-down:
	go tool goose -dir migrations/ postgres $(PG_DSN) down

goose:
	go tool goose -dir migrations/ postgres $(PG_DSN) $(ARGS)

run:
	go run ./cmd/server/main.go

gen-docs:
	go tool swag fmt
	go tool swag init -g internal/controller/rest/server.go -o docs/swagger --v3.1 --generatedTime
	mv docs/swagger/swagger.yaml docs/swagger/openapi.yaml
	mv docs/swagger/swagger.json docs/swagger/openapi.json
	sed -i 's|swag/v2|swag|' docs/swagger/docs.go
	sed -i 's|3.1.0|3.0.1|' docs/swagger/docs.go
	sed -i 's|3.1.0|3.0.1|' docs/swagger/openapi.yaml
	sed -i 's|3.1.0|3.0.1|' docs/swagger/openapi.json

up:
	docker compose up -d

down:
	docker compose down

golangci-lint:
	go tool golangci-lint run --config ./.golangci.yml ./...

k6-get-test:
	k6 run ./loadtest/get-test.js
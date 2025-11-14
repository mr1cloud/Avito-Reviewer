include .env

PG_DSN := "postgresql://$(STORE_USER):$(STORE_PASSWORD)@$(STORE_HOST):$(STORE_PORT)/$(STORE_DATABASE)?sslmode=disable"

migrate-up:
	go tool goose -dir migrations/ postgres $(PG_DSN) up

migrate-down:
	go tool goose -dir migrations/ postgres $(PG_DSN) down

goose:
	go tool goose -dir migrations/ postgres $(PG_DSN) $(ARGS)

run:
	docker compose up -d

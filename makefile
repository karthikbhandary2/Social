include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations/
export DB_ADDR = postgres://Karthik:123@localhost/social?sslmode=disable

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./main.go -d ./cmd/api,./internal && swag fmt

.PHONY: migrate-force
migrate-force:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) force $(ARG)

.PHONY: test
test:
	@go test -v ./...
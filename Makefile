include .env
export

TOOLS_PATH=bin/tools
migrate=$(TOOLS_PATH)/migrate

MIGRATIONS_PATH=migrations

$(migrate):
	GOBIN=`pwd`/$(TOOLS_PATH) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

setup: $(migrate)

.PHONY: run
run:
	go run cmd/main.go

.PHONY: compose.up
compose.up:
	docker compose -f deployments/dev/compose.yml up --build --no-log-prefix --attach question

.PHONY: compose.down
compose.down:
	docker compose down

.PHONY: migrate.create
migrate.create: $(migrate)
	$(migrate) create -ext sql -dir $(MIGRATIONS_PATH) $(name) 

.PHONY: migrate.up
migrate.up: $(migrate)
	$(migrate) -database $(POSTGRES_CONN) -path $(MIGRATIONS_PATH) up

.PHONY: migrate.down
migrate.down: $(migrate)
	$(migrate) -database $(POSTGRES_CONN) -path $(MIGRATIONS_PATH) down


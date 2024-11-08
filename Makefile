include .env
export

TOOLS_PATH=bin/tools
migrate=$(TOOLS_PATH)/migrate
ogen=$(TOOLS_PATH)/ogen

MIGRATIONS_PATH=migrations

$(migrate):
	GOBIN=`pwd`/$(TOOLS_PATH) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

$(ogen):
	GOBIN=`pwd`/$(TOOLS_PATH) go install github.com/ogen-go/ogen/cmd/ogen@v1.6.0

setup: $(migrate) $(ogen)

.PHONY: generate.api
generate.api: $(ogen)
	$(ogen) --loglevel error --clean --config .ogen.yml --target ./api/oas ./api/openapi.yml
	docker run --rm -v `pwd`:/spec redocly/cli bundle ./api/openapi.yml > ./api/bundle.yml


.PHONY: run
run:
	go run cmd/main.go

# COMPOSE
.PHONY: compose.up
compose.up:
	docker compose -f deployments/dev/compose.yml up --build --no-log-prefix --attach question

.PHONY: compose.down
compose.down:
	docker compose down

# MIGRATION
.PHONY: migrate.create
migrate.create: $(migrate)
	$(migrate) create -ext sql -dir $(MIGRATIONS_PATH) $(name) 

.PHONY: migrate.up
migrate.up: $(migrate)
	$(migrate) -database $(POSTGRES_CONN) -path $(MIGRATIONS_PATH) up

.PHONY: migrate.down
migrate.down: $(migrate)
	$(migrate) -database $(POSTGRES_CONN) -path $(MIGRATIONS_PATH) down


# Makefile
.PHONY: all

NETWORK_NAME=statistico_internal

check-network:
	@echo "Checking if network '$(NETWORK_NAME)' exists..."
	@if [ -z "$$(docker network ls --filter name=^$(NETWORK_NAME)$$ --format="{{ .Name }}")" ]; then \
  		echo "Creating network '$(NETWORK_NAME)'..."; \
    	docker network create $(NETWORK_NAME); \
    else \
		echo "Network '$(NETWORK_NAME)' already exists."; \
	fi

docker-build: check-network
	docker compose -f docker-compose.build.yml up -d --build

docker-up: check-network
	docker compose -f docker-compose.build.yml up -d --build --force-recreate

docker-down:
	docker compose -f docker-compose.build.yml down -v

sam-up: check-network
	export CGO_ENABLED=0; \
	sam build --template ./serverless.yaml; \
	sam local start-api --docker-network statistico-football-data_default --env-vars=./env.json

migrate:
	./wait-for-it.sh ${DB_HOST}:${DB_PORT} -t 90 && \
	goose -dir ./database/migrations/ postgres "host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable" up && \
	exit 0

test:
	docker compose -f docker-compose.build.yml run test gotestsum -f short-verbose $(args)

docker-logs:
	docker compose -f docker-compose.build.yml logs -f $(service)

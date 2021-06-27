#!make
NETWORKS="$(shell docker network ls)"
VOLUMES="$(shell docker volume ls)"
POSTGRES_DB="$(shell cat ./secrets/postgres_db)"
POSTGRES_USER="$(shell cat ./secrets/postgres_user)"
POSTGRES_PASSWORD="$(shell cat ./secrets/postgres_passwd)"
SUCCESS=[ done "\xE2\x9C\x94" ]

# default arguments
user ?= root
service ?= api

all: traefik-network postgres-network redis-network postgres-volume
	@echo [ starting client '&' api... ]
	docker-compose up --build traefik api web db redis

traefik-network:
ifeq (,$(findstring traefik-public,$(NETWORKS)))
	@echo [ creating traefik network... ]
	docker network create traefik-public
	@echo $(SUCCESS)
endif

postgres-network: 
ifeq (,$(findstring postgres-net,$(NETWORKS)))
	@echo [ creating postgres network... ]
	docker network create postgres-net
	@echo $(SUCCESS)
endif

redis-network:
ifeq (,$(findstring redis-net,$(NETWORKS)))
	@echo [ creating redis network... ]
	docker network create redis-net
	@echo $(SUCCESS)
endif

postgres-volume:
ifeq (,$(findstring postgres-db,$(VOLUMES)))
	@echo [ creating postgres volume... ]
	docker volume create postgres-db
	@echo $(SUCCESS)
endif

api: traefik-network postgres-network redis-network postgres-volume
	@echo [ starting api... ]
	docker-compose up traefik api db redis

down:
	@echo [ teardown all containers... ]
	docker-compose down
	@echo $(SUCCESS)

tidy:
	@echo [ cleaning up unused $(service) dependencies... ]
	@make exec service="api" cmd="go mod tidy"

exec:
	@echo [ executing $(cmd) in $(service) ]
	docker-compose exec -u $(user) $(service) $(cmd)
	@echo $(SUCCESS)

test-api:
	@echo [ running api tests... ]
	docker-compose run api go test -coverprofile coverage.out -v ./...
	@echo [ outputting coverage.html... ]
	docker-compose run api go tool cover -html=coverage.out -o coverage.html

test-client:
	@echo [ running web client tests... ]
	docker-compose run web yarn test

debug-api:
	@echo [ debugging api... ]
	docker-compose up traefik debug-api db redis

debug-db:
	@# advanced command line interface for postgres
	@# includes auto-completion and syntax highlighting
	@# https://www.pgcli.com/
	@docker run -it --rm --net postgres-net dencold/pgcli postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@db:5432/$(POSTGRES_DB)

dump:
	@echo [ dumping postgres backup for $(POSTGRES_DB)... ]
	@docker exec -it db pg_dump --username $(POSTGRES_USER) $(POSTGRES_DB) > ./api/scripts/backup.sql
	@echo $(SUCCESS)

.PHONY: all
.PHONY: traefik-network
.PHONY: postgres-network
.PHONY: redis-network
.PHONY: postgres-volume
.PHONY: api
.PHONY: down
.PHONY: tidy
.PHONY: exec
.PHONY: test-api
.PHONY: debug-api
.PHONY: debug-db
.PHONY: dump
# Определяем переменные
DOCKER_COMPOSE = docker-compose
SERVICE_NAME = space-vpx-sattelite-dev

# Сообщения
BUILD_MSG = Building the $(SERVICE_NAME) service...
START_MSG = Starting the $(SERVICE_NAME) service...
STOP_MSG = Stopping the $(SERVICE_NAME) service...
REMOVE_MSG = Removing the $(SERVICE_NAME) service...
DB_START_MSG = Starting the PostgreSQL database...
DB_STOP_MSG = Stopping the PostgreSQL database...
DB_REMOVE_MSG = Removing the PostgreSQL database...

# Команды
.PHONY: all build up down restart logs db-start db-stop db-remove

build:
	@echo "$(BUILD_MSG)"
	@docker compose -f ./docker-compose.yml down space-vpx-sattelite-dev
	@docker compose -f ./docker-compose.yml build space-vpx-sattelite-dev
	@docker compose -f ./docker-compose.yml up space-vpx-sattelite-dev -d


up: build
	@echo "$(START_MSG)"
	@$(DOCKER_COMPOSE) up -d
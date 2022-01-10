DOCKER_COMPOSE_FILE_DIR := ./docker/docker-compose.yml
API_CONTAINER_NAME := interior_chair_api
DB_CONTAINER_NAME := interior_chair_mysql

start:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d && \
	docker exec -it $(API_CONTAINER_NAME) go run main.go

restart:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) down && \
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) build --no-cache && \
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d && \
	docker exec -it $(API_CONTAINER_NAME) go run main.go

mysql:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d && \
	docker exec -it $(DB_CONTAINER_NAME) mysql -u root -p

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) down

shell:
	docker exec -it $(API_CONTAINER_NAME) ash

rebuild:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) build --no-cache

tidy:
	docker exec -it $(API_CONTAINER_NAME) go mod tidy

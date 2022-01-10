DOCKER_COMPOSE_FILE_DIR := ./docker/docker-compose.yml
CONTAINER_NAME := interior_chair_api

start:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d && \
	docker exec -it $(CONTAINER_NAME) go run main.go

restart:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) down && \
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) build --no-cache && \
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d && \
	docker exec -it $(CONTAINER_NAME) go run main.go

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) up -d

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) down

shell:
	docker exec -it $(CONTAINER_NAME) ash

rebuild:
	docker-compose -f $(DOCKER_COMPOSE_FILE_DIR) build --no-cache

mysql:
	docker exec -t interior_chair_mysql mysql -u root -p

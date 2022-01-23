API_SERVICE_NAME := app
DB_SERVICE_NAME := db

start:
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) air

restart:
	docker-compose down && \
	docker-compose build --no-cache && \
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) sql-migrate up && \
	docker-compose exec $(API_SERVICE_NAME) air

mysql:
	docker-compose up -d && \
	docker-compose exec $(DB_SERVICE_NAME) mysql -u root -p

shell:
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) ash

tidy:
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) go mod tidy

upd:
	docker-compose up -d

down:
	docker-compose down

rebuild:
	docker-compose build --no-cache

migrate:
	docker-compose exec $(API_SERVICE_NAME) sql-migrate up

migrate-rollback:
	docker-compose exec $(API_SERVICE_NAME) sql-migrate down

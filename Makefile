API_SERVICE_NAME := app
DB_SERVICE_NAME := db

start:
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) air

restart:
	docker-compose down && \
	docker-compose build --no-cache && \
	docker-compose up -d && \
	docker-compose exec $(API_SERVICE_NAME) air

mysql:
	docker-compose up -d && \
	docker-compose exec $(DB_SERVICE_NAME) mysql -u root -p

upd:
	docker-compose up -d

down:
	docker-compose down

shell:
	docker-compose exec $(API_SERVICE_NAME) ash

rebuild:
	docker-compose build --no-cache

tidy:
	docker-compose exec $(API_SERVICE_NAME) go mod tidy

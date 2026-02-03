include .env
export

DB_URL = postgres://postgres:$(DB_PASSWORD)@localhost:5436/postgres?sslmode=disable

up:
	docker compose up -d

down:
	docker compose down

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down

migrate-make:
	migrate create -ext sql -dir ./migrations -seq init $(name)

run:
	air

debug:
	echo $(DB_URL)

start: up migrate-up run
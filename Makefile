build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

migrate:
	docker-compose exec api go run database/migration/migrate.go

seed:
	docker-compose exec api go run database/seeder/seeder.go

test:
	docker-compose exec api go test ./... -cover

swagger:
	docker-compose exec api swag init -g api/main.go

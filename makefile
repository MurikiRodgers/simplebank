postgres:

	docker run -p 5433:5432 --name postgres18 -e POSTGRES_PASSWORD=8901 -d postgres:18.1-bookworm

createdb:

	docker exec -it postgres18 createdb -h localhost -p 5432 -U postgres simple_bank

dropdb:
	docker exec -it postgres18 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:8901@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:8901@localhost:5433/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres createdb migrateup migratedown dropdb sqlc

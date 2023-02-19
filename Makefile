postgres:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres bank

dropdb:
	docker exec -it postgres15 dropdb --username=postgres --owner=postgres bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropbd migrateup migratedown sqlc test

postgres:
	docker run --name dbbank -e POSTGRES_USER=admin1 -e POSTGRES_PASSWORD=admin2 -p 5432:5432 -d postgres
createdb:
	docker exec -it dbbank createdb --username=admin1 --owner=admin1 gobank
dropdb:
	docker exec -it dbbank dropdb -U admin1 gobank
migrateup:
	migrate -path db/migration -database "postgresql://admin1:admin2@localhost:5432/gobank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://admin1:admin2@localhost:5432/gobank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
runserver:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
postgres:
	podman run -d --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:16
createdb:
	podman exec -it postgres createdb --username=root --owner=root simplebank
dropdb:
	podman exec -it postgres dropdb simplebank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/dash-sarthak/simpleBank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock

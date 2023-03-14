postgres:
	docker-compose -f platform/docker-compose.yaml up -d
down:
	docker-compose -f platform/docker-compose.yaml down
createdb:
	docker exec -it simple_bank_postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it simple_bank_postgres dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ku-kim/learn-go-with-project-simplebank/db/sqlc Store
.PHONY: postgres down createdb dropdb migrateup migratedown sqlc server mock
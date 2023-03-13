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

.PHONY: postgres down createdb dropdb migrateup migratedown
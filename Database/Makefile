postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
stopPostgres:
	docker stop postgres14; docker rm postgres14
createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres14 dropdb --username=postgres simple_bank
migrateup:
	.\migrate -verbose -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" up
migratedown:
	.\migrate -verbose -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" down
.PHONE: postgres stopPostgres createdb dropdb migrateup migratedown

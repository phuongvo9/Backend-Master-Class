postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres14 dropdb --username=postgres simple_bank

.PHONE: postgres createdb dropdb
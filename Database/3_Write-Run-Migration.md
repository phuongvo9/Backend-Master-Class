# Write and Run Database Migration
## Table of Contents
1. [Table of Contents](##Table of Contents)
2. [Understand UP/DOWN migration](##Understand UP/DOWN migration)
3. [Explore the manual steps](#Explore the manual steps)
4. [Automate with Makefile](##Automate with Makefile)
5. [Migrate databases](#Migrate databases)
## Install Golang-Migrate
https://github.com/golang-migrate/migrate

`migrate create -ext sql -dir db/migration -seq init_schema`


```aidl
.\migrate create -ext sql -dir db/migration -seq init_schema
C:\Users\Phuong.VoHuy\LEARN\Repos\Backend-Master-Class\Database\db\migration\000001_init_schema.up.sql
C:\Users\Phuong.VoHuy\LEARN\Repos\Backend-Master-Class\Database\db\migration\000001_init_schema.down.sql

```

---
## Understand UP/DOWN migration
![img_2.png](img_2.png)
Up scripts can make FORWARD change to *Database Schema*
DOWN scripts can REVERSE change from UP scripts

### UP scripts
![img_3.png](img_3.png)
Making FORWARD changes to new DB schema *sequentially*
### Down scripts
REVERSING changes from reverse order of UP scripts
![img_4.png](img_4.png)

---
## Init: UP scripts and DOWN scripts
Copy SimpleDB.sql to UP scripts: \Database\db\migration\000001_init_schema.up.sql
Write REVERSE script to DOWN: \Database\db\migration\000001_init_schema.down.sql

## Explore the manual steps - Run Docker container manually 
### Initialize Postgre container
`docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:14-alpine`

### Execute container inside
`docker exec -it postgre14 /bin/sh`


Create a database simple bank inside docker

`createdb --username=postgres --owner=postgres simplebank`

`psql simple_bank -U postgres`

In interative mode container, Login as postgres user and create simple_bank database

`createdb --username=postgres --owner=postgres simple_bank`

`dropdb --username=postgres simple_bank`
### Manually Execute container outside
````
docker exec -it postgres14 'createdb --username=postgres --owner=postgres simple_bank'
docker exec -it postgres14 'psql simple_bank -U postgres'
````

---

## Automate with Makefile
Create a ***Makefile*** with these content:
````
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

````
For example, You can type `make createdb` to quickly create db inside the running container

`make postgres`
````
$ make postgres
docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
748bcce161e0f35d98a4a2c29aabb9100ee2ba7e55ef19ad045ae7da44147cab
````

---

# Migrate databases

``` 
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --versbose up
```


For example
````
$ .\migrate.exe -verbose -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" up
2022/06/04 13:06:38 Start buffering 1/u init_schema
2022/06/04 13:06:38 Read and execute 1/u init_schema
2022/06/04 13:06:38 Finished 1/u init_schema (read 21.626ms, ran 134.5563ms)
2022/06/04 13:06:38 Finished after 167.3983ms
2022/06/04 13:06:38 Closing source and database

---

make migratedown
$ .\migrate.exe -verbose -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" down
2022/06/04 13:09:37 Are you sure you want to apply all down migrations? [y/N]
y
2022/06/04 13:09:42 Applying all down migrations
2022/06/04 13:09:42 Start buffering 1/d init_schema
2022/06/04 13:09:42 Read and execute 1/d init_schema
2022/06/04 13:09:42 Finished 1/d init_schema (read 19.21ms, ran 44.7811ms)
2022/06/04 13:09:42 Finished after 4.8418731s
2022/06/04 13:09:42 Closing source and database
````
Now we use `make` for more convenience
``` 
make migrateup

make migratedown
```
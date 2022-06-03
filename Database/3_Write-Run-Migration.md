# Write and Run Database Migration
## Install Golang-Migrate
https://github.com/golang-migrate/migrate

`migrate create -ext sql -dir db/migration -seq init_schema`


```aidl
.\migrate create -ext sql -dir db/migration -seq init_schema
C:\Users\Phuong.VoHuy\LEARN\Repos\Backend-Master-Class\Database\db\migration\000001_init_schema.up.sql
C:\Users\Phuong.VoHuy\LEARN\Repos\Backend-Master-Class\Database\db\migration\000001_init_schema.down.sql

```
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


## Init: UP scripts and DOWN scripts
Copy SimpleDB.sql to UP scripts: \Database\db\migration\000001_init_schema.up.sql
Write REVERSE script to DOWN: \Database\db\migration\000001_init_schema.down.sql


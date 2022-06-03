# Design Schema

## Database diagram
Go to

Creating tables
````
Table accounts as A {
id bigserial [pk] // auto-increment
owner varchar
balance bigint
currency varchar
created_at timestamptz
}
````
`timestamptz` type helps to track time with the timezone
`bigserial` in Postgres can use at auto-increment value
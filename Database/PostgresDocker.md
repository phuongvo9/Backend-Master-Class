# Work with Postgres container
Check out https://hub.docker.com/_/postgres
## Init: Postgres database with Docker
```
docker images
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```
We choose the specific image ```postgres:14-alpine```
Expose container port 5232 map to host port 5242 ```-p 5432:5432```

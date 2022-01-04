# Golang Unit Testing
Tutorial melakukan unit testing di Golang yang sudah menerapkan clean architecture

## Menjalankan Service
```bash
PSQL_HOST=<IP Database Server> PSQL_PORT=<Port Database Server> PSQL_DBNAME=<Database Name> PSQL_USER=<Database user name> PSQL_PASSWD=<Database User Password> API_HOST=<IP Web Service> API_PORT=<Port Web Service> go run enigmacamp.com/completetesting
```

## Menjalankan Test
```shell
go test -v ./... -coverprofile=cover.out  && go tool cover -html=cover.out
go test -v enigmacamp.com/completetesting/usecase -coverprofile=cover.out  && go tool cover -html=cover.out
```

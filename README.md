# go-rest
rest api

### env variables
- SERVER_PORT
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_DB
- DB_HOST
- GOOSE_DRIVER
- GOOSE_DBSTRING
- GOOSE_MIGRATION_DIR

### added libraries
bash```
go get github.com/labstack/echo/v4
go get github.com/lib/pq
go get golang.org/x/crypto/bcrypt
go install github.com/pressly/goose/v3/cmd/goose@latest
```

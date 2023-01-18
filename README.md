# rotech-post-comment
Simple Rest API (JSON:API) for managing posts and comments

# Introduction

This app manages posts and their comments

# CompileDaemon

CompileDaemon -command="./rotech-post-comment" -build="go build ./cmd/server/"

## Set .env variables

Create .env file in base directory.

```
DATABASE_URL = postgresql://usr:pw@db:5432/postgres
POSTGRES_USER = 
POSTGRES_PASSWORD = 
POSTGRES_DB = rotechpost_db
PORT = 8080
```

## Run Service

```shell
go run ./pkg/migrate/migrate.go
go build ./cmd/server/
./server.exe
```

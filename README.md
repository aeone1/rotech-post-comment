# rotech-post-comment
Simple Rest API (JSON:API) for managing posts and comments

# Introduction

This app manages posts and their comments

## API V1 Path
- POSTS
  - POST - Create a post ```{{server}}/v1/posts```
  - PATCH -  Update a post by id ```{{server}}/v1/posts/:ID```
  - DELETE -  Delete a post by id ```{{server}}/v1/posts/:ID```
  - GET byID - Find a post by id ```{{server}}/v1/posts/:ID```
  - GET count - Count the number of posts that have not been deleted ```{{server}}/v1/posts/count```
  - GET list - Fetch all posts ```{{server}}/v1/posts```

- COMMENTS
  - POST
  - PATCH
  - DELETE
  - GET byID
  - GET list

# CompileDaemon

CompileDaemon -command="./server.exe" -build="go build ./cmd/server/"

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

# REST API example with Go using the Standard Library

This is a simple REST API example built with Go using the standard library. It demonstrates how to create a basic server, handle routes, query a postgres db,  and respond to HTTP requests.

The only external dependency is the PostgreSQL driver for Go from <https://github.com/lib/pq>.

## Installation

```bash
git clone https://github.com/benceluzsinszky/go-stdlib-rest-api.git

cd go-stdlib-rest-api

go get github.com/lib/pq
```

## Usage

The server needs a PostgreSQL database to connect to. You need to export the environment `DATABASE_URL` for the database connection string before running the server:

```bash
export DATABASE_URL="postgres://user:password@localhost/dbname?sslmode=disable"
```

To run the server, use the following command:

```bash

go run . [PORT]
```

The default port is 8080 if not specified.

## Endpoints

| Method   | Endpoint      | Parameters | Description                  |
|----------|---------------|------------|------------------------------|
| `POST`   | `/items`      | `name`     | Create a new item            |
| `GET`    | `/items`      |            | List all items               |
| `GET`    | `/items/{id}` |            | Get a specific item by ID    |
| `PUT`    | `/items/{id}` | `name`     | Update a specific item by ID |
| `DELETE` | `/items/{id}` |            | Delete a specific item by ID |

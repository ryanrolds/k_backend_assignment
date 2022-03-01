# Kong Assignment - Service Domain API

This solution uses a backend REST API written in Go and PostgreSQL DB. The backend services uses Mux for routing and `pqx` for the database driver. I opted to use the `pqx` driver and the `sql.DB` interface because if I got to writing tests I would have used [`sqlmock`](https://github.com/DATA-DOG/go-sqlmock) which uses the `sql.DB` interface, not the native `pqx` interface. The backend service is separated into an API layer, domain objects, service layer, and persistence layer. This approach allows for easier evolution as the endpoint logic is not coupled to the database; If the persistence layer changes, the API layer and services layer will not need to be changed.

Search is provided by PostgreSQL full-text search, which can be seen in `./internal/persistence/service.go`. The DB is prepared using [`migrate`](https://github.com/golang-migrate/migrate), which I'm not that happy with; It really should run the migrations inside of a transaction so that if the migration fails the DB is not placed in an inconsistent state. Other [migration tools](https://bun.uptrace.dev/guide/migrations.html#migration-names) I've wrap migrations in transactions.

I spent about 5 hours on this and copied some boilerplate code/files from other personal projects. If I were to spend more time on this the first thing I would do is introduce [`counterfeiter`](https://github.com/maxbrunsfeld/counterfeiter) and write some unit and contract tests. After that, I would write a data generation script that would populate the DB with at least 100k services and 1M versions. Then I would use [Locust](https://locust.io/) to load test the API, especially the FTS search.

## Setup

Requires:
  * Docker or Docker Desktop - https://docs.docker.com/get-docker/

> The project makes use of Task, which is an alternative to Make. The project will run without [Task](https://taskfile.dev/) installed locally when using Docker. To run without Docker, you will need Go 1.17 or later and Task (or run the commands in Taskfile.yml manually). Not using Task or Docker will require that the env vars in `.env` be set.

## Running with Docker

Start Postgres, run the migrator, and start the remaining services
```
docker-compose up -d postgres
docker-compose up migrator
docker-compose up 
```

At this point the backend service should be available at `http://localhost:8080/`.

## API 

#### Fetch single service and versions by service id
```
// GET /services/{id}
curl -s http://localhost:8080/services/14f777b2-997e-11ec-b909-0242ac120002 | jq .
{
  "status": "ok",
  "total": 1,
  "result": {
    "id": "14f777b2-997e-11ec-b909-0242ac120002",
    "name": "Orders Service",
    "description": "Handles ordering and checkout process",
    "versions": [
      {
        "id": "0.1",
        "service_id": "14f777b2-997e-11ec-b909-0242ac120002",
        "created_at": "2022-03-01T10:43:13.130678-08:00",
        "updated_at": "2022-03-01T10:43:13.130678-08:00"
      },
      {
        "id": "0.2",
        "service_id": "14f777b2-997e-11ec-b909-0242ac120002",
        "created_at": "2022-03-01T10:43:13.130678-08:00",
        "updated_at": "2022-03-01T10:43:13.130678-08:00"
      },
      {
        "id": "0.3",
        "service_id": "14f777b2-997e-11ec-b909-0242ac120002",
        "created_at": "2022-03-01T10:43:13.130678-08:00",
        "updated_at": "2022-03-01T10:43:13.130678-08:00"
      },
      {
        "id": "1.0",
        "service_id": "14f777b2-997e-11ec-b909-0242ac120002",
        "created_at": "2022-03-01T10:43:13.130678-08:00",
        "updated_at": "2022-03-01T10:43:13.130678-08:00"
      }
    ],
    "version_count": 4,
    "created_at": "2020-02-28T09:25:17-08:00",
    "updated_at": "2020-02-28T09:25:18-08:00"
  }
}
```

#### Fetch first page of services

Pagination is supporting using `offset` and `limit` query params. The advantage to using offset vs a page number is that when increasing the limit the client does not loose it's place in the list and simply shows more/less results from the same place in data.

```
// GET /services?offset=0&limit=5
curl -s "http://localhost:8080/services?offset=0&limit=5" | jq .
{
  "status": "ok",
  "total": 50,
  "result": [
    {
      "id": "14f777b2-997e-11ec-b909-0242ac120002",
      "name": "Orders Service",
      "description": "Handles ordering and checkout process",
      "version_count": 4,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:18-08:00"
    },
    {
      "id": "14f79454-997e-11ec-b909-0242ac120002",
      "name": "User Service",
      "description": "Stores user and handles authentication & permissions",
      "version_count": 4,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:19-08:00"
    },
    {
      "id": "14f795f8-997e-11ec-b909-0242ac120002",
      "name": "Team Service",
      "description": "Manages groups of users",
      "version_count": 3,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:20-08:00"
    },
    {
      "id": "14f79828-997e-11ec-b909-0242ac120002",
      "name": "Mobile Presentation Service",
      "description": "Presentation layer for mobile apps",
      "version_count": 1,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:21-08:00"
    },
    {
      "id": "14f799a4-997e-11ec-b909-0242ac120002",
      "name": "Web Presentation Service",
      "description": "Presentation layer for web apps",
      "version_count": 1,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:22-08:00"
    }
  ]
}
```

#### Searching for services

The offset and limit queries are supported. The results are sorted by relevance.

```
// GET /services?q=search_term
 curl -s "http://localhost:8080/services?search=presentation" | jq .
{
  "status": "ok",
  "total": 50,
  "result": [
    {
      "id": "14f79828-997e-11ec-b909-0242ac120002",
      "name": "Mobile Presentation Service",
      "description": "Presentation layer for mobile apps",
      "version_count": 1,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:21-08:00"
    },
    {
      "id": "14f799a4-997e-11ec-b909-0242ac120002",
      "name": "Web Presentation Service",
      "description": "Presentation layer for web apps",
      "version_count": 1,
      "created_at": "2020-02-28T09:25:17-08:00",
      "updated_at": "2020-02-28T09:25:22-08:00"
    }
  ]
}
```

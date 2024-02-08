## BACKEND

This the backend for the Pet-Keeping-Service project. The backend is written in Go and utilizes Postgres as the database for storing data. The backend handles the business logic and provides RESTful APIs for communication with the frontend.

### Prerequisites

- Go 1.21
- Docker && Docker Compose

### Running the server

To run the server, you must first have a PostgreSQL database running. You can do this by running the following command:

```bash
docker-compose up -d
```

Once the database is running, you can run the server by running the following command:

```bash
go run main.go
```

### You can check the endpoints used in the project in the /server/server.go file


### You can check some already init users at the core/init.go file



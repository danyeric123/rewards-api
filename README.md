# Go Backend Template

This is a basic template for a Go backend with PostgreSQL. It includes best practices and boilerplate code to get you started quickly.

To use this template repo, you will need to change all instances of `github.com/danyeric123/backend-service` to your new name `github.com/$USERNAME/$REPO_NAME` where `USERNAME` is your username and `REPO_NAME` is the new name of the repo

## Features

- **Dockerized Setup**: Easily set up and run the application using Docker and Docker Compose.
- **PostgreSQL Integration**: Pre-configured to connect to a PostgreSQL database.
- **Structured Logging**: Uses `logrus` for structured logging.
- **Environment Variables**: Uses environment variables for configuration.
- **Dependency Injection**: Makes the code more testable and modular.

## Project Structure

```markdown
go-backend-template/
├── cmd/
│ └── main.go
├── db/
│ └── db.go
├── handlers/
│ └── handler.go
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── init.sql
├── Makefile
└── README.md
```

## Getting Started

To use this template repo, follow these steps:

1. **Clone the Repository**: Clone this repository to your local machine.
2. **Replace Module Path**: Change all instances of `github.com/danyeric123/backend-service` to your new module path `github.com/$USERNAME/$REPO_NAME` where `USERNAME` is your GitHub username and `REPO_NAME` is the new name of the repository.
3. **Set Up Environment Variables**: Create a `.env` file in the root directory with the following variables:

   ```properties
   POSTGRES_USER=your_postgres_user
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_DB=your_postgres_db
   POSTGRES_HOST=db
   ```

4. **Build and Run**: Use the provided `Makefile` commands to build and run the application.

### Prerequisites

- Docker
- Docker Compose

### Running the Application

**NOTE**: You will need to have an `.env` file for `POSTGRES_DB` `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_HOST` (see [below](#environment-variables))

1. Build the Docker images:

   ```sh
   make build
   ```

2. Run the application:

   ```sh
   make run
   ```

3. Access the application at `http://localhost:8080`.

### Cleaning Up

To stop and remove the Docker containers, run:

```sh
make clean
```

### Commands

1. `make help` - Show this help message
2. `make check_env` - Check if .env file exists
3. `make build` - Build docker containers
4. `make run` - Run docker containers, rebuild if needed
5. `make clean` - Stop and remove docker containers
6. `make fmt` - Format code
7. `make lint` - Lint code

### Database Initialization

The init.sql file contains SQL statements to initialize the database with some boilerplate data. This file is automatically executed when the PostgreSQL container is started.

### Environment Variables

Set the following environment variables in your `.env` file or in your environment:

- `POSTGRES_USER`: The PostgreSQL user
- `POSTGRES_PASSWORD`: The PostgreSQL password
- `POSTGRES_DB`: The PostgreSQL database name
- `POSTGRES_HOST`: The PostgreSQL host (usually db when using Docker Compose)

### Explanation of Chosen Packages

- **logrus**: `logrus` is used for structured logging. It provides log levels, hooks, and formatters, making it a powerful and flexible logging library compared to the standard `log` package.
- **mux**: `mux` is a powerful URL router and dispatcher for matching incoming requests to their respective handler. Though the standard library's `http` package would suffice, `mux` adds more flexibility and features such as variables in routes, middleware support, and more.
- **sqlx**: `sqlx` is an extension of the standard `database/sql` package. It has become a standard for working with SQL databases in Go due to its additional functionalities.
- **pq**: `pq` is a pure Go Postgres driver for the `database/sql` package. PostgreSQL is a powerful, open-source object-relational database system, and `pq` provides reliable and efficient connectivity to PostgreSQL databases.

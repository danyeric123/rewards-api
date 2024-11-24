# Receipt Processor

This project is a receipt processor that includes the endpoints `POST /receipts/process` and `GET /receipt/{id}/points`. Additionally, tests have been added to facilitate easier testing of different scenarios.

My inclination was to try and incorporate DDD principles in the code since this was something quite important to several repos I have worked in and I have found it quite helpful.

**NOTE**: This project builds upon a template I previously developed (see [danyeric123/backend-service](https://github.com/danyeric123/backend-service)). While the template provided a foundation, most of the code is newly written specifically for this project.


## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Background](#background)

## Features

- **Dockerized Setup**: Easily set up and run the application using Docker and Docker Compose.
- **PostgreSQL Integration**: Pre-configured to connect to a PostgreSQL database.
- **Structured Logging**: Uses `logrus` for structured logging.
- **Environment Variables**: Uses environment variables for configuration.
- **Dependency Injection**: Makes the code more testable and modular.
- **GORM Integration**: Uses GORM for ORM, making it easier to work with the database.

## Project Structure

```markdown
receipt-processor/
├── cmd/
│   └── main.go
├── db/
│   ├── config.go
│   ├── models.go
│   ├── receipt.go
│   └── receipt_test.go
├── domain/
│   ├── receipt.go
│   └── receipt_test.go
├── handlers/
│   └── receipt.go
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Getting Started

To use this template repo, follow these steps:

1. **Clone the Repository**: Clone this repository to your local machine.username and `REPO_NAME` is the new name of the repository.
2. **Set Up Environment Variables**: Create a `.env` file in the root directory with the following variables:

   ```properties
   POSTGRES_USER=your_postgres_user
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_DB=your_postgres_db
   POSTGRES_HOST=db
   ```

3. **Build and Run**: Use the provided `Makefile` commands to build and run the application.

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
3. `make check_network` - Check if Docker network exists and create it if not
4. `make build` - Build Docker containers
5. `make run` - Run Docker containers, rebuild if needed
6. `make clean` - Stop and remove Docker containers
7. `make fmt` - Format code
8. `make lint` - Lint code
9. `make test` - Run tests
10. `make test_package` - Run tests in a specific package

### Environment Variables

Set the following environment variables in your `.env` file or in your environment:

- `POSTGRES_USER`: The PostgreSQL user
- `POSTGRES_PASSWORD`: The PostgreSQL password
- `POSTGRES_DB`: The PostgreSQL database name
- `POSTGRES_HOST`: The PostgreSQL host (usually db when using Docker Compose)

## Background

### Choice of Language

I work with Golang for several projects and given `FetchRewards` uses it, I saw it as a great opportunity to practice my skills in setting up a repository from scratch.

### Explanation of Chosen Packages

- **logrus**: `logrus` is used for structured logging. It provides log levels, hooks, and formatters, making it a powerful and flexible logging library compared to the standard `log` package.
- **mux**: `mux` is a powerful URL router and dispatcher for matching incoming requests to their respective handler. Though the standard library's `http` package would suffice, `mux` appears to be the standard.
  - As I wrote it, I thought to use `fiber` or `chi` but I felt like there was too much custom features that might not be ideal within the Golang world, which values a more Unix philosophy
- **gorm**: `gorm` is an ORM library for Go. It simplifies database interactions and allows you to think in terms of Domain-Driven Design (DDD).

### Specific Choices

Some specific choices have comments on assumptions or how things might be done differently with more time.

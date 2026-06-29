# Go React Router

A backend service written in Go, focused on building scalable, maintainable, and testable REST APIs using Clean Architecture principles.

The project emphasizes clear separation of concerns by isolating domain logic from infrastructure, allowing business rules to evolve independently of frameworks, databases, and delivery mechanisms.


## Architecture

The application follows a layered architecture.

```text
HTTP Request
      │
      ▼
   Handler
      │
      ▼
   Use Case
      │
      ▼
 Repository
      │
      ▼
 PostgreSQL
```

### API Layer

Responsible for transport concerns.

* HTTP routing
* Request binding
* Response serialization
* HTTP status codes

### Domain Layer

Contains the application's business logic.

* Entities
* Use Cases
* Repository contracts
* Domain errors
* Business validation

The domain does not depend on any infrastructure packages.

### Platform Layer

Contains infrastructure implementations.

Examples include:

* PostgreSQL
* Password hashing
* External services
* Storage
* Email providers

Implementations satisfy interfaces defined by the domain.


## Project Structure

```text
cmd/
└── server/

internal/
├── api/
│   ├── handler/
│   ├── response/
│   └── routes/
│
├── app/
│
├── config/
│
├── domain/
│   └── user/
│       ├── usecase/
│       ├── errors.go
│       ├── repository.go
│       └── user.go
│
└── platform/
    ├── database/
    └── security/
```


## Design Principles

* Feature-oriented organization
* Dependency inversion
* Interface-driven architecture
* Constructor-based dependency injection
* Explicit application bootstrap
* Framework-independent domain layer
* Thin HTTP handlers
* Repository pattern
* Stateless use cases


## Technologies

* Go
* Gin
* PostgreSQL
* pgx
* bcrypt


## Configuration

Configuration is loaded from environment variables.

Example:

```env
PORT=8080

DATABASE_CONNECTION_STRING=postgresql://...

JWT_SECRET=...

JWT_EXPIRY=24h
```

## Running the Project

Install dependencies:

```bash
go mod download
```

Run the application:

```bash
go run ./cmd/server
```


## Development Workflow

Application startup performs the following steps:

1. Load configuration
2. Initialize database connection
3. Configure router
4. Register routes
5. Start HTTP server

Each feature follows the same execution flow:

```text
Route
    ↓
Handler
    ↓
Use Case
    ↓
Repository
    ↓
Database
```


## Goals

This project aims to serve as a production-oriented backend demonstrating:

* Clean Architecture
* Domain-driven design principles
* Scalable project organization
* Dependency injection
* Maintainable API design
* Testable business logic
* Clear separation between domain and infrastructure

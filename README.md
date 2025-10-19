# Easy Attend Service

Backend service for the Easy Attend application built with Go, Gin, and PostgreSQL.

## Project Structure

```
easy-attend-service/
├── main.go                      # Application entry point
├── cmd/
│   └── migrate/
│       └── main.go              # Database migration command
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   ├── database.go          # Database connection
│   │   └── migrate.go           # Database migrations
│   ├── handlers/
│   │   └── health.go            # HTTP handlers
│   ├── models/
│   │   └── models.go            # Database models
│   ├── routes/
│   │   └── routes.go            # Route definitions
│   └── services/                # Business logic
├── pkg/
│   └── utils/                   # Utility functions
├── api/
│   └── v1/                      # API documentation
├── migrations/                  # Database migrations
├── .env                         # Environment variables
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd easy-attend-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment variables by copying `.env.example` to `.env`:
```bash
cp .env.example .env
```

4. Update the `.env` file with your database credentials:
```
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_DATABASE=easy-attend
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=your_password
SERVER_PORT=8080
```

## Running the Application

1. Make sure your PostgreSQL database is running

## Running the Application

1. Make sure your PostgreSQL database is running

2. Run database migration (first time only):
```bash
cd cmd/migrate
go run main.go
```

If tables already exist, you can force migration:
```bash
cd cmd/migrate
go run main.go --force
```

3. Run the application:
```bash
go run main.go
```

The server will start on the port specified in your `.env` file (default: 8080).

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check if the service is running
- `GET /api/v1/version` - Get API version information

## Development

### Adding New Features

1. Create models in `internal/models/`
2. Add handlers in `internal/handlers/`
3. Define routes in `internal/routes/`
4. Implement business logic in `internal/services/`

### Database Migrations

Database migrations will be stored in the `migrations/` directory.

## Technologies Used

- **Go** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **godotenv** - Environment variable management
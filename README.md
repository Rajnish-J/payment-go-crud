# Payment Go CRUD

A REST API for payment management system built with Go, Gin, and GORM.

## Project Overview

This is a clean architecture Go application that handles:
- **User Management** - User registration, authentication
- **Company Management** - Company CRUD operations  
- **Product Management** - Product catalog management
- **Payment Processing** - Payment history and order processing

## Architecture

```
├─ cmd/           # Application entry point
├─ config/        # Database configuration
├─ controller/    # HTTP handlers
├─ service/       # Business logic layer
├─ repository/    # Data access layer
├─ model/         # Database models
├─ dto/           # Data transfer objects
├─ routes/        # Route definitions
├─ business/      # Validation logic
└─ utils/         # Helper utilities
```

## Required Packages

### Core Dependencies
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
```

### Swagger Documentation
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

## Quick Start

1. **Clone & Install**
   ```bash
   git clone <repository>
   cd payment-go-crud
   go mod tidy
   ```

2. **Setup Environment**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Generate Swagger Docs**
   ```bash
   swag init -g src/cmd/main.go
   ```

4. **Run Application**
   ```bash
   go run src/cmd/main.go
   ```

## API Documentation
Visit: http://localhost:8080/swagger/index.html
# Golang Project Template

## Specification
- Init Project using Golang [Cobra](https://github.com/spf13/cobra)
- Parse and Get Env Configuration using [Viper](https://github.com/spf13/viper)
- Implement Http Server Router [Echo](https://echo.labstack.com/docs)
- Connect to MySQL Database Management System via [Gorm](https://gorm.io/docs/)
- Prepare local testing and start up with [Docker Compose](https://docs.docker.com/manuals/)

## Dependencies

### Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.4
```

### Cobra
```bash
go install github.com/spf13/cobra-cli@latest
```

# Directory Layout
This project uses Monolithic Modular Architecture.
```
── cmd
├── database
├── deployment
│   └── local
│       └── data
├── docs
│   └── swagger
├── internal
│   ├── contract
│   ├── handler
│   │   ├── authentication
│   │   ├── health
│   │   └── profile
│   ├── middleware
│   ├── model
│   ├── registry
│   │   ├── authentication
│   │   ├── health
│   │   └── profile
│   ├── repository
│   │   └── user
│   └── service
│       ├── authentication
│       └── profile
├── migrations
│   ├── data
│   │   └── versions
│   └── schema
│       └── versions
├── server
├── static
└── util
    └── hashing
```

## Swagger
* Play url:
```
http://localhost:3000/swagger/index.html
```
* Generate swagger specification
```bash
swag init --parseDependency --parseDependencyLevel 3 -g main.go -g handler.go -d ./internal/handler -o ./docs/swagger
```

## Server Startup
1. Edit `local.env` file to specific environment configuration
2. Make sure the database service is open and running
3. At the root directory, Run the following command
    ```bash
    go run main.go serve
    ```
4. Migrate database schema and data
    ```bash
    go run main.go migration migrate --schema --data
    ```
# JWT Authentication with Go

This is a basic JWT (JSON Web Token) authentication project using Gin framework (Golang) and PostgreSQL.

## API Endpoints

| Method | Endpoint  | Description                  |
| ------ | --------- | ---------------------------- |
| POST   | /signup   | Register new user            |
| POST   | /login    | Login and generate JWT token |
| GET    | /validate | Validate JWT token           |

## Installation Steps

1. Install Go from [official website](https://golang.org/dl/)

2. Initialize project and install dependencies:

```bash
go mod init jwt-authentication-go
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4
go get github.com/joho/godotenv
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

3. Create `.env` file:

```env
PORT=3000
DB_URL="host=localhost user=postgres password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
SECRET="your-secret-key"
```

## Running the Application

Standard run:

```bash
go run main.go
```

With hot-reload:

```bash
CompileDaemon -command="./jwt-authentication-go"
```

Test the APIs using Postman or any API testing tool of your choice.

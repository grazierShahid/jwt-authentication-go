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

2. Clone the repository:

```bash
git clone https://github.com/grazierShahid/jwt-authentication-go.git
cd jwt-authentication-go
```

3. Initialize project and install dependencies:

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

4. Create `.env` file:

```env
PORT=3000
DB="host=localhost user=postgres password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
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

## 🚀 API Documentation

### 🔐 Authentication Endpoints

Our JWT authentication system provides secure, stateless authentication with the following endpoints:

### 1️⃣ User Registration

Create a new user account with email verification and secure password hashing.

```http
POST /signup
```

**Request Body:**

```json
{
  "email": "test@example.com",
  "password": "pass123",
  "name": "Mr. Test User"
}
```

**Successful Response** `HTTP 200`

```json
{
  "msg": "Signup successfull!"
}
```

**Error Cases:**

- `400 Bad Request`: Invalid input data
- `400 Bad Request`: Password hashing failed
- `400 Bad Request`: User creation failed

### 2️⃣ User Authentication

Authenticate user and receive a JWT token for secured API access.

```http
POST /login
```

**Request Body:**

```json
{
  "email": "test@example.com",
  "password": "pass123"
}
```

**Successful Response** `HTTP 200`

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

✨ **Features:**

- JWT token valid for 30 days
- Secure cookie setting with HTTP-only flag
- Protection against XSS and CSRF attacks

**Error Cases:**

- `400 Bad Request`: Invalid credentials
- `400 Bad Request`: Authentication failed
- `400 Bad Request`: Token generation failed

### 3️⃣ Token Validation

Verify your authentication status and get user information.

```http
GET /validate
```

**Authentication:**
The endpoint uses the HTTP-only cookie set during login. No additional headers required.

**Successful Response** `HTTP 200`

```json
{
  "user info": {
    "ID": 1,
    "Email": "test@example.com",
    "Name": "Mr. Test User"
  }
}
```

## 🔒 Auth Flow

1. **Login:**

   - Get JWT token in response
   - Auto-set secure HTTP-only cookie
   - 30 days validity

2. **Protected Routes:**
   - Cookies handle auth automatically
   - Same-site Lax protection
   - No manual token needed

## 🛠️ Quick Test Guide

1. **Login:** `POST /login`

   - Note down token (optional)
   - Cookie sets automatically

2. **Validate:** `GET /validate`

   - Just call endpoint
   - Cookie handles auth

3. **Test Security:**
   - Clear cookies = logout
   - 30 days expiry

## 🔒 Security

- Bcrypt hashing
- JWT tokens
- HTTP-only cookies
- Same-site protection
- 30-day expiration

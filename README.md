# 🧑‍💼 Go + Gin + PostgreSQL

REST API built with Golang, Gin framework, and PostgreSQL. Supports full CRUD operations for:

- Users
- User Profiles
- Leave Balances

## 🚀 Features

- CRUD operations on users
- Related data:
  - Profile (national ID, birth place, etc.)
  - Leave balances (annual, sick, unpaid)
- PostgreSQL integration
- Docker & Docker Compose ready
- Environment-based configuration

---

## 🧱 Project Structure

├── docker-compose.yml
├── Dockerfile
├── .env
├── main.go
├── database
│  └── database.go
├── services
│  └── user_service.go
├── models
│  ├── user.go
│  ├── profile.go
│  └── leave_balance.go
├── validators
│  └── validator.go
└── handlers
    └── user.go

## 🔧 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## ⚙️ Setup

1. **Clone the repository:**

```bash
git clone https://github.com/yrvvan/golang-project.git
cd golang-project
```
2. Create .env file
```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hris
APP_PORT=8089
```
3. Start the application
```
docker-compose up --build
```
It will :
- Start PostgreSQL
- Run the Go backend on http://localhost:8089

4. Endpoints
- **GET	/users** #Get all users
- **GET	/users/:id** #Get user by ID
- **POST /users**	#Create a user
- **PUT	/users/:id** #Update user
- **DELETE /users/:id**	#Delete user

## 🛠️ Tech Stack
1. Go (Golang)
2. Gin Web Framework
3. GORM (ORM for Go)
4. PostgreSQL
5. Docker & Docker Compose

## Docker Compose

```javascript
version: "3.9"

services:
  app:
    build: .
    container_name: go-rose
    ports:
      - "8089:8089"
    env_file:
      - .env
    depends_on:
      - db

  db:
    image: postgres:17-alpine
    container_name: postgres-db
    ports:
      - "5432:5432"
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: postgres
    volumes:
      - postgres-data:/var/lib/postgresql/data

volumes:
  postgres-data:
```

## 🧹 To Do
1. Add JWT authentication
2. Role-based access control
3. CI/CD pipeline
4. Frontend integration (React?)

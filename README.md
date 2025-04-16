# 🧑‍💼 Go + Gin + PostgreSQL

REST API built with Golang.

Supports full CRUD operations for:
- Users
- User Profiles
- Leave Balances

## 🚀 Features

- CRUD operations on users
- Related data:
  - Profile (national ID, birth place, etc.)
  - Leave balances (annual, sick, unpaid)
- PostgreSQL integration
- JWT Authorization
- Password Hashing
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
JWT_SECRET={set by yourself}
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

| Method  | Path  | Purpose |
| :------------ |:---------------:| -----:|
| GET      | /users | Get all users |
| GET      | /users/:id        |   Get user by id |
| POST | /users        |    Create a user |
| PUT | /users/:id        |    Update a user |
| DELETE | /users/:id        |    Delete a user |

## 🛠️ Tech Stack
1. Go (Golang)
2. Gin Web Framework
3. GORM (ORM for Go)
4. PostgreSQL
5. JWT Authorization
6. Middleware
7. Password Hashing
8. Docker & Docker Compose

## Usage / Example
![Screenshot 2025-04-12 134451](https://github.com/user-attachments/assets/9810ccff-bd4b-4242-a4d6-d18741b4b224)
![Screenshot 2025-04-12 134423](https://github.com/user-attachments/assets/f003b496-b344-465b-8e01-d25271f06eb6)

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
1. Role-based access control
2. Send Logs into Grafana
3. CI/CD pipeline
4. Frontend integration (React?)

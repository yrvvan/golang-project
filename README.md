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
|  └── database.go
├── services
|  └── user_service.go
├── models
|  ├── user.go
|  ├── profile.go
│  └── leave_balance.go
├── validators
│  └── validator.go
├── handlers
   └── user.go

## 🔧 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## ⚙️ Setup

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/hris-golang.git
cd hris-golang

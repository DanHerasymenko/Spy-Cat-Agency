# Spy Cat Agency API

**Author:** [Daniil Herasymenko](https://github.com/DanHerasymenko)

**Default swagger link:** [http://localhost:8082/swagger/index.html#/](http://localhost:8082/swagger/index.html#/)

**Default PostgreSQL connection URL (internal Docker network):**  
`postgres://spycat:spycat@postgres:5432/spycat?sslmode=disable`


This application implements a management system for the Spy Cat Agency, where cat agents are assigned espionage missions with specific targets. The system supports full CRUD operations for spy cats, missions, and targets, enforces business logic rules, and integrates with external services like TheCatAPI to validate cat breeds.

---

##  Technologies Used

- Golang 1.24.0
- PostgreSQL 17.4
- Web Framework: Gin
- Migrations: Goose
- Swagger: swaggo/gin-swagger
- Docker + Docker Compose
- External API: [TheCatAPI](https://thecatapi.com)

---

##  Run with Docker

1. Clone the repository.
2. Create a `.env` file in the root directory using the template below.
3. Put your TheCatAPI key in `CAT_API_KEY` in the `.env` file.
3. Start the application using Docker Compose:

```bash
docker compose up --build
```

---

## ️ Example `.env` File

```env
# Application
APP_PORT=:8080

# PostgreSQL
DB_HOST=postgres
DB_PORT=5432
DB_USER=spycat
DB_PASS=spycat
DB_NAME=spycat

# TheCatAPI
CAT_API_URL=https://api.thecatapi.com/v1
CAT_API_KEY=your_key_here

# Docker Compose
APP_CONTAINER_PORT=8080
APP_LOCAL_PORT=8082
POSTGRES_CONTAINER_HOST=postgres
POSTGRES_LOCAL_PORT=5435
```

---

##  Migrations

- Migrations run automatically on application start via `migrations` service in docker-compose.

---

##  Implemented Endpoints

### 🐱 Cats

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | `/api/cats/create`        | Create a new spy cat |
| GET    | `/api/cats/list`          | List all spy cats |
| GET    | `/api/cats/{id}`          | Get a spy cat |
| PUT    | `/api/cats/{id}/salary`   | Update cat salary |
| DELETE | `/api/cats/{id}`          | Delete a spy cat |

---

### 🎯 Missions & Targets

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/missions`                     | List all missions |
| POST   | `/api/missions`                     | Create a mission |
| GET    | `/api/missions/{id}`                | Get mission by ID |
| PUT    | `/api/missions/{id}`                | Update mission (e.g., mark as completed) |
| DELETE | `/api/missions/{id}`                | Delete mission (only if unassigned) |
| POST   | `/api/missions/{id}/assign`         | Assign a cat to a mission |
| POST   | `/api/missions/{id}/targets`        | Add a target (if < 3 & mission not completed) |
| PUT    | `/api/missions/targets/{id}`        | Update a target (notes, completed flag) |
| DELETE | `/api/missions/targets/{id}`        | Delete a target (only if not completed) |

---

##  Example Create Cat JSON

```json
{
   "breed": "Bambino",
   "name": "Murzik",
   "salary": 300,
   "years_experience": 5
}
```

---

##  Example Create Mission JSON

```json
{
  "cat_id": 3,
  "name": "Operation Name",
  "targets": [
    {
       "country": "Ukraine",
       "name": "Some name",
      "notes": "Some notes"
    }
  ]
}
```

---


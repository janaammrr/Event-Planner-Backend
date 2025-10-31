# Event Planner Backend (Go + Gin + MySQL)

This is a minimal backend for an Event Planner app with Signup and Login using bcrypt and JWT.

- Gin framework
- GORM with MySQL driver (connection is optional; falls back to in-memory store)
- Env via `.env`
- CORS enabled for `http://localhost:4200`

## Folder Structure

- `config/`: environment and database setup
- `controllers/`: request handlers (signup/login)
- `models/`: GORM models
- `routes/`: router and endpoints
- `utils/`: helpers (bcrypt, JWT, response helpers)

## Prerequisites

- Go 1.21+

## Run the backend

1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. (Optional) Create `.env` from the example:
   ```bash
   cp env.example .env
   ```
3. Start the server:
   ```bash
   go run main.go
   ```

Server runs on `http://localhost:8080`.

If database environment variables are not set, the app runs with an in-memory user store so you can try signup/login immediately.

## Connect to MySQL later

Your schema creates a database named `EventPlanner`. Use that name for `DB_NAME`.

1. Ensure a MySQL database exists and is reachable.
2. Set the following in `.env`:
   ```env
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_NAME=EventPlanner
   JWT_SECRET=your_strong_secret
   ```
3. Start the server. On startup the app will:
   - Establish the GORM connection
   - Auto-migrate the `User` model

If any required DB env var is missing or connection fails, the app logs a message and continues without DB (using in-memory store).

## API Endpoints

Base path: `/api`

### Health

- `GET /api/ping`
  - Response: `{"message":"pong"}`

### Signup

- `POST /api/signup`
- Request JSON:
  ```json
  {
    "name": "Alice",
    "email": "alice@example.com",
    "password": "secret123"
  }
  ```
- Success 201:
  ```json
  {
    "id": 1,
    "name": "Alice",
    "email": "alice@example.com",
    "role": "attendee"
  }
  ```
- Errors: 400 invalid payload, 409 email already registered, 500 server error

### Login

- `POST /api/login`
- Request JSON:
  ```json
  {
    "email": "alice@example.com",
    "password": "secret123"
  }
  ```
- Success 200:
  ```json
  {
    "token": "<jwt>",
    "user": {
      "id": 1,
      "name": "Alice",
      "email": "alice@example.com",
      "role": "attendee"
    }
  }
  ```
- Errors: 400 invalid payload, 401 invalid email or password, 500 server error

## Notes

- Passwords are hashed with bcrypt.
- JWTs are signed using HS256 with `JWT_SECRET`.
- CORS allows origin `http://localhost:4200`.
- Models align with `event_planer_DB/event_planer_schema.sql`:
  - Table `users`, PK `user_id` (INT UNSIGNED, auto-increment)
  - Columns and constraints: `name`, `email` (unique `ux_users_email`), `password_hash`, `role` enum with default `attendee`, `created_at` (indexed `ix_users_created_at`).

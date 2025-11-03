# Go To-Do List (Go + SQLite)

Teaching-ready backend demonstrating:
- SQL-first design (DDL in `sql/schema.sql`, DML in `sql/seed.sql`)
- REST endpoints:
  - `GET /api/todos`
  - `POST /api/todos` (non-idempotent)
  - `PUT /api/todos/:id` (idempotent)

## Run
```bash
go run .

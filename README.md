<p align="center">
  <img src="https://github.com/Thee-Hector-Genaro-Pacheco/go-to-do-list/blob/main/G0-TO-DO_API.png" alt="Go To-Do API Banner" width="100%">
</p>
<p align="center">
  <em>Teaching REST, SQL, and GraphQL Fundamentals — Built with Go + SQLite</em>
</p>
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

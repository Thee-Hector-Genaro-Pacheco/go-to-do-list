<p align="center">
  <img src="https://github.com/Thee-Hector-Genaro-Pacheco/go-to-do-list/blob/main/G0-TO-DO_API.png" alt="Go To-Do API Banner" width="100%">
</p>
<p align="center">
  <em>Teaching REST, SQL, and GraphQL Fundamentals — Built with Go + SQLite</em>
</p>

<p align="center">
  <a href="https://go.dev/"><img alt="Go" src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white"></a>
  <a href="https://www.sqlite.org/"><img alt="SQLite" src="https://img.shields.io/badge/DB-SQLite-003B57?logo=sqlite&logoColor=white"></a>
  <a href="#"><img alt="REST" src="https://img.shields.io/badge/API-REST-1f6feb"></a>
  <a href="#"><img alt="GraphQL Ready" src="https://img.shields.io/badge/GraphQL-ready-DF019A?logo=graphql&logoColor=white"></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg"></a>
  <a href="https://github.com/Thee-Hector-Genaro-Pacheco/go-to-do-list/issues"><img alt="PRs welcome" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>
</p>

# Go To-Do List (Go + SQLite)

This repo is a **teaching-ready** backend that demonstrates:
- **SQL-first design** (DDL & DML) with a lightweight **SQLite** database
- **REST fundamentals**, including **POST vs PUT (idempotency)**
- A clean path to **GraphQL** (coming next) using the same database

---

## 🌐 Endpoints

| Method | Path                  | Purpose                      |
|-------:|-----------------------|------------------------------|
| GET    | `/api/todos`          | List all todos               |
| POST   | `/api/todos`          | Create a new todo            |
| PUT    | `/api/todos/:id`      | Replace/update a todo by id  |

> **Idempotency note:**  
> `POST` **creates** a new resource each time (non-idempotent).  
> `PUT /:id` **replaces** the resource at `:id` with the same body every time (idempotent).

---

## ▶️ Run locally

```bash
go run .
# server listens on http://localhost:8080

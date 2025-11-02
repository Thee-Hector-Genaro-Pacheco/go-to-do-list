package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func restRoutes(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// GET /api/todos
	r.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY id DESC`)
		if err != nil { http.Error(w, err.Error(), 500); return }
		defer rows.Close()
		var list []Todo
		for rows.Next() {
			var t Todo
			var c int
			if err := rows.Scan(&t.ID, &t.Title, &c, &t.CreatedAt, &t.UpdatedAt); err != nil {
				http.Error(w, err.Error(), 500); return
			}
			t.Completed = c == 1
			list = append(list, t)
		}
		writeJSON(w, 200, list)
	}).Methods("GET")

	// POST /api/todos — create (NON-idempotent)
	r.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
			http.Error(w, "invalid body", 400); return
		}
		res, err := db.Exec(`INSERT INTO todos (title, completed) VALUES (?, ?)`, in.Title, boolToInt(in.Completed))
		if err != nil { http.Error(w, err.Error(), 500); return }
		id, _ := res.LastInsertId()
		t, err := getTodo(db, id)
		if err != nil { http.Error(w, err.Error(), 500); return }
		writeJSON(w, 201, t)
	}).Methods("POST")

	// PUT /api/todos/{id} — replace (IDEMPOTENT)
	r.HandleFunc("/api/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if id <= 0 { http.Error(w, "invalid id", 400); return }

		var in struct {
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
			http.Error(w, "invalid body", 400); return
		}
		res, err := db.Exec(`UPDATE todos SET title=?, completed=? WHERE id=?`, in.Title, boolToInt(in.Completed), id)
		if err != nil { http.Error(w, err.Error(), 500); return }
		aff, _ := res.RowsAffected()
		if aff == 0 { http.Error(w, "not found", 404); return }

		t, err := getTodo(db, id)
		if err != nil { http.Error(w, err.Error(), 500); return }
		writeJSON(w, 200, t)
	}).Methods("PUT")

	return r
}

func getTodo(db *sql.DB, id int64) (*Todo, error) {
	var t Todo
	var c int
	err := db.QueryRow(`SELECT id, title, completed, created_at, updated_at FROM todos WHERE id=?`, id).
		Scan(&t.ID, &t.Title, &c, &t.CreatedAt, &t.UpdatedAt)
	if err != nil { return nil, err }
	t.Completed = c == 1
	return &t, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func boolToInt(b bool) int {
	if b { return 1 }
	return 0
}

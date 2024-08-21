package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateUserController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form.Get("nama")
			email := r.Form.Get("email")
			umur := r.Form.Get("umur")
			gender := r.Form.Get("gender")
			_, err := db.Exec("INSERT INTO karyawan (nama, email, umur, gender) VALUES (?, ?, ?, ?)", name, email, umur, gender)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create_user.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

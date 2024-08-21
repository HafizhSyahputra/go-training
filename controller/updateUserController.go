package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateUserController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
		
			id_karyawan := r.Form.Get("id_karyawan")  
			name := r.Form.Get("nama")
			email := r.Form.Get("email")
			umur := r.Form.Get("umur")
			gender := r.Form.Get("gender")
		
			_, err := db.Exec("UPDATE karyawan SET nama = ?, email = ?, umur = ?, gender = ? WHERE id_karyawan = ?", name, email, umur, gender, id_karyawan)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}else if r.Method == "GET" {
			id := r.URL.Query().Get("id_karyawan")

			row := db.QueryRow("SELECT nama, email, umur, gender FROM karyawan WHERE id_karyawan = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var karyawann Karyawan
			err := row.Scan(
				&karyawann.Nama,
				&karyawann.Email,
				&karyawann.Umur,
				&karyawann.Gender,
			)
			karyawann.IdKaryawan = id
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update_user.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			data := make(map[string]any)
			data["karyawann"] = karyawann
			

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

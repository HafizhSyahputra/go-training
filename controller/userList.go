package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Karyawan struct {
	IdKaryawan string
	Nama       string
	Email      string
	Umur       int
	Gender     string
}

func NewUserList(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_karyawan, nama, email, umur, gender FROM karyawan")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var karyawans []Karyawan
		for rows.Next() {
			var karyawann Karyawan

			err = rows.Scan(
				&karyawann.IdKaryawan,
				&karyawann.Nama,
				&karyawann.Email,
				&karyawann.Umur,
				&karyawann.Gender,
			)
			

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			karyawans = append(karyawans, karyawann)
		}

		fp := filepath.Join("views", "user_list.html")
		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["karyawans"] = karyawans

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

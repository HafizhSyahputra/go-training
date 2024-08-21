package controller

import (
	"database/sql"
	"fmt"
	"net/http"
)

func DeleteUserController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            id := r.FormValue("id_karyawan")

            fmt.Println("ID Karyawan yang ingin dihapus:", id)
            
            result, err := db.Exec("DELETE FROM karyawan WHERE id_karyawan = ?", id)
            if err != nil {
                w.Write([]byte(err.Error()))
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            rowsAffected, _ := result.RowsAffected()
            fmt.Println("Rows Affected:", rowsAffected)
            
            if rowsAffected == 0 {
                w.Write([]byte("Data tidak ditemukan atau tidak ada yang dihapus"))
                w.WriteHeader(http.StatusNotFound)
                return
            }

            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
    }
}



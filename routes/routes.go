package routes

import (
	"database/sql"
	"net/http"

	"github.com/HafizhSyahputra/go-training/controller"
)

func MapRoute(server *http.ServeMux, db*sql.DB){
	server.HandleFunc("/", controller.NewUserList(db))
	server.HandleFunc("/userList/create", controller.CreateUserController(db))
	server.HandleFunc("/userList/update", controller.UpdateUserController(db))

}
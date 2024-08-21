package main

import (
	"net/http"

	"github.com/HafizhSyahputra/go-training/database"
	"github.com/HafizhSyahputra/go-training/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoute(server, db)

	http.ListenAndServe(":8080", server)

}

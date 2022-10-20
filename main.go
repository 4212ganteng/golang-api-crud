package main

import (
	"fmt"
	"golang-api/database"
	"golang-api/pkg/mysql"
	"golang-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.Konekdb()

	database.RunMigratation()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000",r)
}
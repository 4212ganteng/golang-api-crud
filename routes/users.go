package routes

import (
	"golang-api/handlers"
	"golang-api/pkg/mysql"
	"golang-api/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
UserRepository := repositories.RepositoryUser(mysql.DB)
h := handlers.HandlerUser(UserRepository)


r.HandleFunc("/users", h.FindUsers).Methods("GET")
r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
r.HandleFunc("/users/create", h.CreateUser).Methods("POST") 
r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PATCH")
r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}
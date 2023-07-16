package main

import (
	"net/http"

	"github.com/HiramZednem/go-crud-mysql/db"
	"github.com/HiramZednem/go-crud-mysql/models"
	"github.com/HiramZednem/go-crud-mysql/routes"
	"github.com/gorilla/mux"
)



func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	

	r := mux.NewRouter()

	
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

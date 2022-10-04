package main

import (
	//"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/mnarvaezm96/go-gorm-restapi/db"
	"github.com/mnarvaezm96/go-gorm-restapi/models"
	"github.com/mnarvaezm96/go-gorm-restapi/routes"
)

func main() {

	db.DBconnection()

	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3003", r)
}

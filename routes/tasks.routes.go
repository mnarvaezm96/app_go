package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mnarvaezm96/go-gorm-restapi/db"
	"github.com/mnarvaezm96/go-gorm-restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return

	}

	json.NewEncoder(w).Encode(&task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	db.DB.Find(&task)
	json.NewEncoder(w).Encode(&task)

	if task.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return

	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) //204
}

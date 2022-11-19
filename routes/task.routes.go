package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/go-restapi/config"
	"github.com/osmait/go-restapi/model"
)

func GetTasksHandle(w http.ResponseWriter, r *http.Request) {
	var tasks []model.Task
	config.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)

}

func GetTaskHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task model.Task
	config.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
	}
	json.NewEncoder(w).Encode(&task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	json.NewDecoder(r.Body).Decode(&task)
	createdTask := config.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task model.Task
	config.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task nt found"))
		return
	}

	config.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)

}

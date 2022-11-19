package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/go-restapi/config"
	"github.com/osmait/go-restapi/model"
)

func GetUsersHandle(w http.ResponseWriter, r *http.Request) {
	var user []model.User
	config.DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUserHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user model.User
	config.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not found"))
		return
	}
	config.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandle(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := config.DB.Create(&user)
	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user model.User

	config.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("User not found"))
		return
	}
	config.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

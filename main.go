package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/go-restapi/config"
	"github.com/osmait/go-restapi/model"
	"github.com/osmait/go-restapi/routes"
)

func main() {
	r := mux.NewRouter()
	config.DBConnetion()

	config.DB.AutoMigrate(model.User{})
	config.DB.AutoMigrate(model.Task{})

	// Users Routes
	r.HandleFunc("/user", routes.GetUsersHandle).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandle).Methods("GET")
	r.HandleFunc("/user", routes.PostUserHandle).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DETELE")

	// Task Routes

	r.HandleFunc("/task", routes.GetTasksHandle).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandle).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DETELE")

	http.ListenAndServe(":3000", r)
}

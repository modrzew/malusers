package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
	"github.com/modrzew/malusers/api"
)

func addHandlers(router *mux.Router, db *gorm.DB) {
	handlers := &api.Handlers{DB: db}
	router.HandleFunc("/user/{username}", handlers.GetUserStats).Methods("GET")
}

func main() {
	db := malusers.OpenDb()

	router := mux.NewRouter()
	addHandlers(router, db)
	http.ListenAndServe("localhost:8000", router)
}

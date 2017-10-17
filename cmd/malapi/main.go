package main

import (
	"fmt"
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
	config := malusers.ReadConfig()
	db := malusers.OpenDb()

	router := mux.NewRouter()
	addHandlers(router, db)
	url := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)
	http.ListenAndServe(url, router)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
	"github.com/modrzew/malusers/api"
)

func logger(inner http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func addHandlers(router *mux.Router, db *gorm.DB, cache *api.Cache) {
	handlers := &api.Handlers{DB: db, Cache: cache}
	router.Handle("/user/{username}", logger(handlers.GetUserStats)).Methods("GET")
}

func main() {
	config := malusers.ReadConfig()
	db := malusers.OpenDb()
	cache := api.GetCache(db)

	router := mux.NewRouter()
	addHandlers(router, db, cache)
	url := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)
	http.ListenAndServe(url, router)
}

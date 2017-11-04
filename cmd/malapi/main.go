package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/api"
	"github.com/modrzew/malusers/core"
)

func logger(inner http.HandlerFunc) http.HandlerFunc {
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

func cors(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		inner(w, r)
	})
}

func addHandlers(router *mux.Router, db *gorm.DB, cache *api.Cache) {
	handlers := &api.Handlers{DB: db, Cache: cache}
	router.Handle("/user/{username}", logger(cors(handlers.GetUserStats))).Methods("GET")
	router.Handle("/stats/{kind:(?:anime|manga)}/{group:(?:gender|year)}", logger(cors(handlers.GetGlobalStats))).Methods("GET")
}

func main() {
	config := core.ReadConfig()
	db := core.OpenDb()
	cache := api.GetCache(db)

	router := mux.NewRouter()
	addHandlers(router, db, cache)
	url := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)
	http.ListenAndServe(url, router)
}

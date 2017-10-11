package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
)

// DBHandler wraps mux handlers and provides them access to DB
type DBHandler struct {
	Handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request)
	DB      *gorm.DB
}

// Serve executes handler method and passes handler to DB to it
func (h *DBHandler) Serve(w http.ResponseWriter, r *http.Request) {
	h.Handler(h.DB, w, r)
}

// GetUserStats returns JSON info about single user
func GetUserStats(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := malusers.User{}
	db.Where("username = ?", params["username"]).Find(&user)
	if db.NewRecord(user) {
		w.WriteHeader(404)
		return
	}
	stats := UserStats{
		Username: user.Username,
		Birthday: user.Birthday.Time.Format("2006-02-01"),
	}
	json.NewEncoder(w).Encode(stats)
}

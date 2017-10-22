package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
)

// Handlers contains reference to the database and all handlers
type Handlers struct {
	DB    *gorm.DB
	Cache *Cache
}

// GetUserStats returns JSON info about single user
func (h *Handlers) GetUserStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := strings.ToLower(params["username"])
	stats, error := h.Cache.GetUser(username)
	if error != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 not found")
		return
	}
	json.NewEncoder(w).Encode(stats)
}

// GetGlobalStats returns JSON info about global stats
func (h *Handlers) GetGlobalStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kind := strings.ToLower(params["kind"])
	group := strings.ToLower(params["group"])
	if kind == "anime" {
		stats := malusers.GetAnimeStats(h.DB, group)
		json.NewEncoder(w).Encode(stats)
	}
}

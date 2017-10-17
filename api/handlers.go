package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
)

// Handlers contains reference to the database and all handlers
type Handlers struct {
	DB *gorm.DB
}

// GetUserStats returns JSON info about single user
func (h *Handlers) GetUserStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := malusers.User{}
	username := strings.ToLower(params["username"])
	h.DB.Where("username = ?", username).First(&user)
	if h.DB.NewRecord(user) {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 not found")
		return
	}
	h.DB.Where("username = ?", username).Find(&user.AnimeStats)
	h.DB.Where("username = ?", username).Find(&user.MangaStats)
	age := time.Since(user.Birthday.Time).Hours() / 24 / 365
	stats := UserStats{
		Username:   user.DisplayName,
		Age:        int(age),
		LastUpdate: user.UpdatedAt.UTC(),
		AnimeStats: user.AnimeStats,
		MangaStats: user.MangaStats,
	}
	json.NewEncoder(w).Encode(stats)
}
